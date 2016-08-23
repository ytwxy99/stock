# -*-coding: utf-8 -*-

def readfile(filepath):
    """
    Read file.

    :param filepath: file path.
    """
    try:
        fp = open(filepath)
        lines = fp.readlines()[4:]
        contents = list()
        for line in lines:
            contents.append([l for l in line.split('"') if l != ' ' and l != "," and  l != '\r\n' and l != ''])
        return contents
    except Exception as e:
        print e

def get_prices(datas):
    """
    Get prices.

    :param datas: stocks data.
    """
    high_price = 0.0
    low_price = 0.0

    for data in datas:
        if high_price == 0.0 and low_price == 0.0:
            high_price = float(data[2])
            low_price = float(data[2])

        elif float(data[2]) > high_price:
            high_price = float(data[2])

        elif float(data[2]) < low_price:
            low_price = float(data[2])

    return high_price, low_price

def get_vol(datas):
    """
    Get vol.

    :param datas: stocks data.
    """
    high_vol = 0.0
    low_vol = 0.0

    for data in datas:
        if high_vol == 0.0 and low_vol == 0.0:
            high_vol = float(data[1])
            low_vol = float(data[1])

        elif float(data[1]) > high_vol:
            high_vol = float(data[1])

        elif float(data[1]) < low_vol:
            low_vol = float(data[1])

    return high_vol, low_vol

if __name__ == '__main__':
    readfile("000007.csv")
