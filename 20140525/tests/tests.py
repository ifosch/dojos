#!/usr/bin/env python2
import unittest
import math

class TPV_impl:
    products = {}

    def price(self, s):
        if isinstance(s, list) : # Si es un array
            temp = float(0)
            for element in s:
                temp += self.products[element]["price"]
                if self.hasdiscount(element):
                    temp -= self.products[element]["discount"]
        else:
            temp = float(self.products[s]["price"])
            if self.hasdiscount(s):
                temp = round(temp - self.products[s]["discount"], 2)
            
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
        self.impl.products = {
                "A": {"price":2.30, "discount":0.10},
                "B": {"price":1.60}}

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

if __name__ == '__main__':
    unittest.main()
