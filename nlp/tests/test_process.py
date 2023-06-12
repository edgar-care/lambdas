import unittest
from handlers import process
from handlers import request

class TestProcess(unittest.TestCase):
    '''TestProcess'''
    
    def test_process_present(self):
        '''
        Test process function, as to return a context object with maux_de_ventre as true
        '''
        data = request.Req(
            symptoms=['maux_de_ventre'], input='oui')
        self.assertEqual(process.process(data), {
            'context': [{'symptom': 'maux_de_ventre', 'present': True }]})

    def test_process_absent(self):
        '''
        Test process function, as to return a context object with maux_de_ventre as false
        '''
        data = request.Req(symptoms=['maux_de_ventre'], input='non')
        self.assertEqual(process.process(data), {
            'context': [{'symptom': 'maux_de_ventre', 'present': False }] })

    def test_process_unknown(self):
        '''
        Test process function, as to return a context object with maux_de_ventre as None
        '''
        data = request.Req(symptoms=['maux_de_ventre'], input="oue c'est greg")
        self.assertEqual(process.process(data), {'context': [{'symptom': 'maux_de_ventre', 'present': None }]})

    def test_process_new_symptom(self):
        '''
        Test process function, as to return a context object with maux_de_ventre as true
        '''
        data = request.Req(symptoms=[], input="oue c'est le fils de greg et j'ai mal au ventre")
        self.assertEqual(process.process(data), {'context': [{'symptom': 'maux_de_ventre', 'present': True }]})
