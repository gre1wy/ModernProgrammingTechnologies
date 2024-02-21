#З використанням рекурсії розв’язати наступні задачі.
#1. Піднести до додатного цілого степеня дійсне ненульове число.
def power(x, n):
    if x <= 0:
        raise ValueError("x should be positive")
    if n < 0:
        raise ValueError("n should be non-negative")
    # Базовий випадок: x^0 = 1
    if n == 0:
        return 1
    # Рекурсивний випадок: x^n = x * x^(n-1)
    else:
        return x * power(x, n-1)
    
import unittest

class TestPowerFunction(unittest.TestCase):
    def test_power_with_positive_exponent(self):
        self.assertEqual(power(2, 3), 8)
        self.assertEqual(power(5, 4), 625)
        self.assertEqual(power(3, 5), 243)

    def test_power_with_zero_exponent(self):
        self.assertEqual(power(10, 0), 1)
        self.assertEqual(power(7, 0), 1)
        self.assertEqual(power(1, 0), 1)

    def test_power_with_non_positive_base(self):
        with self.assertRaises(ValueError):
            power(0, 3)

        with self.assertRaises(ValueError):
            power(-2, 4)
                
    def test_power_with_negative_base_and_even_exponent(self):
        with self.assertRaises(ValueError):
            power(-3, 2)

print("something t43")
if __name__ == '__main__':
    unittest.main()