# -*- coding: utf-8 -*-

#import matplotlib.pyplot as plt
#fig = plt.figure()

import pylab as pl
import readfile

def draw_table_vol():
    pl.ioff()
    contents = readfile.readfile("000007.csv")
    high, low = readfile.get_prices(contents)
    ranges = range(len(contents))
    prices = [ content[2] for content in contents ]

    table_x = list()
    table_y = list()
    high, low = readfile.get_vol(contents)

    for x, content in enumerate(contents):
        table_x.append(x)
        if float(content[1]) < 10:
            table_y.append(float(content[1]))
        else:
            table_y.append(0)

    pl.plot(table_x, table_y, '-')
    pl.show()

if __name__ == '__main__':
    draw_table_vol()
