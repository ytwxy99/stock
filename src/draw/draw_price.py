# -*- coding: utf-8 -*-

#import matplotlib.pyplot as plt
#fig = plt.figure()

import pylab as pl
import readfile

def draw_table_price():
    pl.ioff()
    contents = readfile.readfile("000007.csv")
    high, low = readfile.get_prices(contents)
    ranges = range(len(contents))
    prices = [ content[2] for content in contents ]

    pl.plot(ranges, prices, color="red")
    #pl.plot([(float(low) - float(low)*0.02), (float(high) + float(high)*0.02)])
    pl.plot([float(low), float(high)])
    pl.title("Hello world")

    pl.show()

if __name__ == '__main__':
    draw_table_price()
