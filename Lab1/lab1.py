import unittest


# З використанням рекурсії розв’язати наступні задачі.
# 1. Піднести до додатного цілого степеня ціле ненульове число.
def power(x, n):

    if x <= 0 or not isinstance(x, int):
        raise ValueError("x should be positive integer")
    if n < 0 or not isinstance(n, int):
        raise ValueError("n should be non-negative integer")

    # Базовий випадок: x^0 = 1
    if n == 0:
        return 1
    # Рекурсивний випадок: x^n = x * x^(n-1)
    else:
        return x * power(x, n-1)


class TestPowerFunction(unittest.TestCase):
    def test_power_with_positive_exponent(self):
        self.assertEqual(power(2, 3), 8)
        self.assertEqual(power(5, 4), 625)
        self.assertEqual(power(3, 5), 243)

    def test_power_with_zero_exponent(self):
        self.assertEqual(power(10, 0), 1)
        self.assertEqual(power(7, 0), 1)
        self.assertEqual(power(1, 0), 1)

    def test_power_with_not_positive_not_integer_base(self):
        with self.assertRaisesRegex(ValueError, 
                                    "x should be positive integer"):
            power(0, 3)
        with self.assertRaisesRegex(ValueError, 
                                    "x should be positive integer"):
            power(-2, 4)
        with self.assertRaisesRegex(ValueError, 
                                    "x should be positive integer"):
            power(1.5, 4)
    def test_power_with_not_integer_or_negative_base(self):
        with self.assertRaisesRegex(ValueError, 
                                    "n should be non-negative integer"):
            power(1, -3)
        with self.assertRaisesRegex(ValueError, 
                                    "n should be non-negative integer"):
            power(2, 1.5)


if __name__ == '__main__':
    unittest.main()

