# -*- coding: utf-8 -*-

import os
import sys
import ftplib


class FTPSync(object):
	def __init__(self):
		self.host = 'ftp.gushenbbs.com'
		self.user = 'gushenbbs170508'
		self.password = 'gushenbbs170508'
		self.local_dir = '/Users/tracy/workspace/stock/src/data'
		self.conn = ftplib.FTP(self.host, self.user, self.password)
		#远端FTP目录
		self.conn.cwd('/') 
		#本地下载目录
		os.chdir(self.local_dir)

	def get_dirs_files(self):
		""" 得到当前目录和文件, 放入dir_res列表"""
		dir_res = []
		files = list()
		self.conn.dir('.', dir_res.append)
		files_tmp = [f for f in dir_res if 'csv' in f]
		print files_tmp
		for file in files_tmp:
			files.append([f for f in file.split(' ') if f != '' ][3])
		dirs = list()
		for dir in dir_res:
			try:
				dirs.append([d for d in dir.split(' ') if d != '' and ('.' not in d)][3])
			except:
				pass
		print files
		return (files, dirs)

	def walk(self, next_dir):
		print 'Walking to', next_dir
		self.conn.cwd(next_dir)
		try:
			os.mkdir(next_dir)
		except OSError:
			pass
		os.chdir(next_dir)

		ftp_curr_dir = self.conn.pwd()
		local_curr_dir = os.getcwd()

		files, dirs = self.get_dirs_files()
		print "FILES: ", files
		print "DIRS: ", dirs
		for f in files:
			print next_dir, ':', f
			outf = open(f, 'wb')
			try:
				self.conn.retrbinary('RETR %s' % f, outf.write)
			finally:
				outf.close()
		for d in dirs:
			os.chdir(local_curr_dir)
			self.conn.cwd(ftp_curr_dir)
			self.walk(d)

	def run(self):
		self.walk('.')

def main():
    f = FTPSync()
    f.run()

if __name__ == '__main__':
	main()
