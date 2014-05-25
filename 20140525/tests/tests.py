#!/usr/bin/env python2
import unittest

class Offers:
    OfferList = {}

    def addOfferPercent(self, name, value):
        self.OfferList[name] = { "%": value }

    def addOfferValue(self, name, value):
        self.OfferList[name] = { "val": value }

    def addOfferAmount(self, name, amount, accounted):
        self.OfferList[name] = (amount, accounted)

    def getOffer(self, name):
        return self.OfferList[name]

class Products:

    ProductList = {}

    def __init__(self, offers):
        self.offers = offers


    def addProduct(self, name, price, offer_name):
        self.ProductList[name] = (price, offer_name)

    def getPrice(self, name, quantity):
         price, offer = self.ProductList[name]
         tmp = self.offers.getOffer(offer)
         if isinstance(tmp, tuple):
             amount, accounted = tmp
             if quantity % amount == 0:
                 return quantity/amount * price * accounted
         else:
             if tmp[1] == "val":
                 return (price - tmp[2]) * quantity

             else:
                 return price * (1 - tmp[2]) * quantity

class TPV_impl:
    products = {}
    plist = {}

    def price(self, s):
        if isinstance(s, list) : # Si es un array
            for element in s:
                if self.plist[element]: 
                    self.plist[element] += 1
                else:
                    self.plist[element] = 1

            temp = float(0)
            for i in self.plist:
                 temp += self.product.getPrice(i, self.plist[i])
        else:
            temp = self.product.getPrice(s, 1)

        return temp


    def hasdiscount(self, element):
        if "discount" in self.products[element]:
            return True
        else:
            return False

    def verifyallprice (self):
        for element in self.products: 
            if "price" not in self.products[element]:
                return False
        return True

class TPV (unittest.TestCase):
    def setUp (self):
        self.impl=TPV_impl()
#        self.impl.products = {
#                "A": {"price":2.30, "discount":0.10},
#                "B": {"price":1.60},
#		"C": {"price":3.60}
#		}

        self.offers = Offers()
        self.products = Products(self.offers)

        self.offers.addOfferAmount("of1", 2, 1)
        self.offers.addOfferAmount("of2", 3, 2)
        self.offers.addOfferValue("of3", -30)
        self.offers.addOfferPercent("of4", 0.10)

        self.products.addProduct("A", 2.30, "of1")
        self.products.addProduct("B", 1.60, "of3")
        self.products.addProduct("C", 3.60, "of4")

    def testFirst(self):
        self.assertEqual(2.20,self.impl.price("A"))

    def testSecond(self):
        self.assertEqual(1.60, self.impl.price("B"))

    def testAddition(self):
        self.assertEqual(3.80, self.impl.price(["A","B"]))

    def testThird(self):
        self.assertTrue(self.impl.hasdiscount("A"))
        self.assertFalse(self.impl.hasdiscount("B"))

    def testFour(self):
        self.assertTrue(self.impl.verifyallprice())

    def testFive(self):
        self.assertEqual(3.90, self.impl.price(["A", "A", "B"]))

    def testSumTotal(self):
        self.assertEqual(6.84, self.impl.price(["A","A","B","C"]))

if __name__ == '__main__':
    unittest.main()
