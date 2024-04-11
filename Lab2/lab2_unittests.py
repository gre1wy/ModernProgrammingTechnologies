import unittest
import lab2 as l
import sys
import io

class TestBubbleSort(unittest.TestCase):

    # testing function bubblesort()
    def test_int(self):
        nums_int = [2, 1, 3, 1, 6, 12, 3, 1]
        expected_result_int = [1, 1, 1, 2, 3, 3, 6, 12]
        self.assertEqual(l.bubble_sort(nums_int), expected_result_int)

    def test_float(self):
        nums_float = [3.5, 1.1, 4.2, 1.7, 5.9, 2.6]
        expected_result_float = [1.1, 1.7, 2.6, 3.5, 4.2, 5.9]
        self.assertEqual(l.bubble_sort(nums_float), expected_result_float)

    def test_int_float(self):
        nums_float = [3, 1, 4.2, 1.7, 5.9, 2.6]
        expected_result_float = [1, 1.7, 2.6, 3, 4.2, 5.9]
        self.assertEqual(l.bubble_sort(nums_float), expected_result_float)

    def test_empty_list(self):
        nums_empty = []
        expected_result_empty = []
        self.assertEqual(l.bubble_sort(nums_empty), expected_result_empty)

    # testing exit code
    def test_exit_code_0(self):
        nums_exit_code_0 = [2, 3, 1, 2, 1, 3.1]
        exit_code_0 = l.main(nums_exit_code_0)[0]
        self.assertEqual(exit_code_0, 0)  

    def test_exit_code_1(self):
        nums_exit_code_1 = ['k', 3, 1, '7', 1, 3.1]
        exit_code_1 = l.main(nums_exit_code_1)
        self.assertEqual(exit_code_1, 1)
        nums_exit_code_1_str = "2 1 1.2 5 4.6"
        exit_code_1_str = l.main(nums_exit_code_1_str)
        self.assertEqual(exit_code_1_str, 1)

    # testing stdin
    def test_stdin(self):
        test_cases = [
            ("2 1 1.2 5 4.6", [1, 1.2, 2, 4.6, 5]), # test_int_float
            ("5.2 3.8 1.6 4.3", [1.6, 3.8, 4.3, 5.2]), # test_float
            ("100 50 200", [50, 100, 200]) # test_int
        ]

        for test_input, expected_result in test_cases:
            with self.subTest(test_input=test_input, expected_result=expected_result):
                sys.stdin = io.StringIO(test_input)
                exit_code, sorted_numbers = l.main()
                self.assertEqual(exit_code, 0)
                self.assertEqual(sorted_numbers, expected_result)
    

if __name__ == '__main__':
    unittest.main()
