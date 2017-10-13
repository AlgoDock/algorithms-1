#! /usr/local/bin/python
import argparse
import os
import pprint
import sqlite3
from collections import namedtuple

import requests
from tqdm import tqdm

pp = pprint.PrettyPrinter()
conn = sqlite3.connect("leetcode.db")
cur = conn.cursor()
table_name = 'problem'
cur.execute("CREATE TABLE IF NOT EXISTS `{}` "
            "(id integer(5) NOT NULL PRIMARY KEY, "
            "title varchar(128), title_slug varchar(128), difficulty tinyint(1), status varchar(16) DEFAULT NULL,  go tinyint(1), "  # NOQA E501
            "updated_at date DEFAULT (strftime('%Y-%m-%d', 'now', 'localtime')))".format(table_name))  # NOQA E501
Prombem = namedtuple('Problem', ['id', 'title', 'title_slug', 'difficulty', 'status', 'go', 'date'])  # NOQA E501


class LeetCode:
    LOGIN_URI = 'https://leetcode.com/accounts/login/'
    HEADERS = {
        'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8',  # NOQA E501
        'Accept-Encoding': 'gzip, deflate, sdch',
        'Accept-Language': 'en-US,en;q=0.8,zh-CN;q=0.6,zh;q=0.4,zh-TW;q=0.2',
        'Connection': 'keep-alive',
        'Host': 'leetcode.com',
        'Referer': 'https://leetcode.com/accounts/login/',
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:56.0) Gecko/20100101 Firefox/56.0', # NOQA E501
        'Content-Type': 'application/x-www-form-urlencoded',
    }

    def __init__(self, username, password, difficulty):
        self.username = username
        self.password = password
        self.difficulty = difficulty
        self.session = requests.Session()
        self.login()

    def _get_csrfmiddlewaretoken(self):
        resp = self.session.get(self.LOGIN_URI, headers=self.HEADERS)
        # <input type='hidden' name='csrfmiddlewaretoken' value='B5yVhk7grKZU4rt3hNxGZ9NUB67lmIW0RuAg75Vx9f3WA8grGIV4qI2eFbmlK6Dq' />  # NOQA E501
        return resp.text.split('csrfmiddlewaretoken')[1].split("'")[2]

    def login(self):
        data = {
            'csrfmiddlewaretoken': self._get_csrfmiddlewaretoken(),
            'login': self.username,
            'password': self.password,
        }
        resp = self.session.post(
            self.LOGIN_URI, data=data, headers=self.HEADERS)
        if resp.status_code != 200:
            print("Login failed.", resp.status_code)
            os._exit(1)

    def all(self):
        """
        profile data
        """
        url = 'https://leetcode.com/api/progress/all/'
        resp = self.session.get(url)
        if resp.status_code != 200:
            print("request failed:", resp.status_code)
            os._exit(1)
        data = resp.json()
        print('{unsolved} Todo | {solvedTotal}/{questionTotal} Solved'
              ' | {attempted} Attempted'.format(**data))
        print('Easy {Easy} Medium {Medium} Hard {Hard}'.format(
            **data['solvedPerDifficulty']))

    def _get_support_languages_by_tilte_slug(self, tilte_slug):
        url = 'https://leetcode.com/problems/{}/description/'.format(
            tilte_slug)
        text = self.session.get(url).text
        codeDefinition = text.split('codeDefinition: ')[1].split(
            'enableTestMode:')[0].strip().strip(',')
        return ([i.split("'")[-2]
                for i in codeDefinition.split(',') if "'text'" in i])

    def sync_all_problems(self):
        url = 'https://leetcode.com/api/problems/algorithms/'
        data = self.session.get(url).json()

        retrieve_sql = 'SELECT * FROM `{table_name}` WHERE id={pk}'
        insert_sql = 'INSERT INTO `{table_name}` (id, title, title_slug, difficulty, status, go) VALUES ({id}, "{title}", "{title_slug}", {difficulty}, "{status}", {go})'  # NOQA E501
        update_sql = 'UPDATE `{table_name}` SET go={go}, status="{status}" WHERE id={pk}'  # NOQA E501

        pairs = sorted(data['stat_status_pairs'],
                       key=lambda x: x['stat']['question_id'])
        pairs = filter(lambda x: x['paid_only'] is False, pairs)
        if self.difficulty:
            pairs = filter(
                lambda x: x['difficulty']['level'] == self.difficulty, pairs)
        for p in tqdm(pairs):
            problem_id = p['stat']['question_id']
            title_slug = p['stat']['question__title_slug']
            status = p['status']

            instance = cur.execute(retrieve_sql.format(
                table_name=table_name, pk=problem_id)).fetchone()
            if instance is None:
                go = 1 if 'Go' in self._get_support_languages_by_tilte_slug(
                    title_slug) else 0
                values = {
                    'id': problem_id,
                    'title': p['stat']['question__title'],
                    'title_slug': title_slug,
                    'difficulty': p['difficulty']['level'],
                    'status': status,
                    'go': go,
                }
                cur.execute(insert_sql.format(table_name=table_name, **values))
                conn.commit()
                continue
            if status is None:
                problem = Prombem(*instance)
                go = problem.go
                if go != 1 and \
                        'Go' in self._get_support_languages_by_tilte_slug(
                            title_slug):
                    go = 1
                cur.execute(update_sql.format(
                    table_name=table_name, go=go, pk=problem_id,
                    status=status))
                conn.commit()

    def get_unsolved_problem(self):
        # get a problem by difficulty filter
        self.sync_all_problems()
        all_sql = 'SELECT * FROM `{table_name}` ORDER BY id'.format(
            table_name=table_name)
        pairs = []
        for p in cur.execute(all_sql).fetchall():
            problem = Prombem(*p)
            if problem.status != 'ac' and problem.go and (
                    self.difficulty is None or
                    problem.difficulty == self.difficulty):
                pairs.append(problem._asdict())
        pp.pprint(pairs[0])
        print('https://leetcode.com/problems/{}/description/'.format(
            pairs[0]['title_slug']))

    def analysis(self, allstatistics, query):
        if allstatistics:
            self.all()
        if query:
            self.get_unsolved_problem()


parser = argparse.ArgumentParser(
    description='Analysis leetcode problems progress.')
parser.add_argument('-u', '--username', type=str, required=True,
                    help='Username for login.')
parser.add_argument('-p', '--password', type=str, required=True,
                    help='Password to use when connecting to leetcode.')
parser.add_argument('-d', '--difficulty', type=int,
                    choices=(1, 2, 3),
                    help='Difficulty of problems.')

parser.add_argument('-a', '--allstatistics', action='store_true',
                    help='Show all statistical information.')
parser.add_argument('-query', '--query', action='store_true',
                    help='Query one unsolved problem order by id.')


if __name__ == '__main__':
    args = parser.parse_args()
    leetcode = LeetCode(username=args.username, password=args.password,
                        difficulty=args.difficulty)
    leetcode.analysis(allstatistics=args.allstatistics, query=args.query)
