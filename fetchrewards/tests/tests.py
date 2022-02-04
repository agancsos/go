#!/usr/bin/env python3
###############################################################################
# Name         : hackerone.py                                                 #
# Author       : Abel Gancsos                                                 #
# Version      : v. 1.0.0.0                                                   #
# Description  : Helps build out Hacker1 reports without a REST API.          #
###############################################################################
import os, sys, requests;

class TestSuite:
	base_endpoint=None;
	def __init__(self, params=dict()):
		self.base_endpoint = params["-b"] if "-b" in params.keys() else "http://localhost:4441";
	def add_credit(self, payer, points, timestamp):
		return requests.post("{0}/credit".format(self.base_endpoint), json={"payer":payer, "points":points, "timestamp":timestamp});
	def use_credit(self, points):
		return requests.post("{0}/spend".format(self.base_endpoint), json={"points":points});
	def invoke(self):
		assert self.base_endpoint != "", "Base endpoint cannot be empty...";
		assert self.add_credit("DANNON", 1000, "2020-11-02T14:00:00Z").status_code == 200, "Add credit failed";
		assert self.add_credit("UNILEVER", 200, "2020-10-31T11:00:00Z").status_code == 200, "Add credit failed";
		assert self.add_credit("DANNON", -200, "2020-10-31T15:00:00Z").status_code != 200, "Add credit did not fail";
		assert self.add_credit("MILLER COORS", 10000, "2020-11-01T14:00:00Z").status_code == 200, "Add credit failed";
		assert self.add_credit("DANNON", 300, "2020-10-31T10:00:00Z").status_code == 200, "Add credit failed";
		rsp = self.use_credit(5000);
		print(rsp.json());
		rsp = requests.get("{0}/balance".format(self.base_endpoint));
		print(rsp.json());
	pass;

if __name__ == "__main__":
	params = dict();
	for i in range(0, len(sys.argv) - 1): params[sys.argv[i]] = sys.argv[i + 1];
	session = TestSuite(params);
	session.invoke();

