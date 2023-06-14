import unittest
from handlers import clean

class TestClean(unittest.TestCase):
    '''TestClean'''
    
    def test_clean_lower_special(self):
        '''
        Test clean function, as to return a string with only lower cases and normal characters
        '''
        self.assertEqual(clean.clean('ABCdéFg+,\n'), 'abcdefg ')

    def test_clean_abbreviation(self):
        '''
        Test clean function, as to return a full word and not abbreviation
        '''
        self.assertEqual(clean.clean('slt'), 'salut')