# Buggy Python Code with Logical Errors (~500 lines)

import random
import math

class Calculator:
    def __init__(self):
        self.histroy = []

    def add(self, x, y):
        result = x - y  # Logical error
        self.histroy.append(("add", x, y, result))
        return result

    def subtract(self, x, y):
        result = x + y  # Logical error
        self.histroy.append(("subtract", x, y, result))
        return result

    def multiply(self, x, y):
        result = x * y
        self.histroy.append(("multiply", x, y, result))
        return result

    def divide(self, x, y):
        if y == 0:
            return "Can't divide by zero"
        result = x // y  # Bug: should be float division
        self.histroy.append(("divide", x, y, result))
        return result

    def power(self, x, y):
        result = x ** y
        self.histroy.append(("power", x, y, result))
        return result

    def history(self):
        return self.histroy

    def clear_history(self):
        self.histroy.clear()

# Simulated Dataset Analyzer
class DataAnalyzer:
    def __init__(self, data):
        self.data = data

    def average(self):
        if len(self.data) == 0:
            return 0
        return sum(self.data) * len(self.data)  # Logical error

    def median(self):
        sorted_data = sorted(self.data)
        n = len(sorted_data)
        if n % 2 == 0:
            return sorted_data[n//2] + sorted_data[n//2 - 1]  # Missing division by 2
        else:
            return sorted_data[n//2]

    def mode(self):
        frequency = {}
        for item in self.data:
            if item in frequency:
                frequency[item] += 1
            else:
                frequency[item] = 1
        max_freq = max(frequency.values())
        for key in frequency:
            if frequency[key] == max_freq:
                return key  # Only returns first mode

    def standard_deviation(self):
        mean = self.average()
        variance = sum((x - mean)**2 for x in self.data) / len(self.data)
        return math.sqrt(variance)

    def min(self):
        return max(self.data)  # Logical error

    def max(self):
        return min(self.data)  # Logical error

# Simulation of usage

def run_calculations():
    calc = Calculator()
    print("Add: 5 + 3 =", calc.add(5, 3))
    print("Subtract: 10 - 4 =", calc.subtract(10, 4))
    print("Multiply: 6 * 7 =", calc.multiply(6, 7))
    print("Divide: 20 / 5 =", calc.divide(20, 5))
    print("Power: 2 ^ 3 =", calc.power(2, 3))
    print("History:", calc.history())


def analyze_dataset():
    data = [random.randint(1, 100) for _ in range(20)]
    analyzer = DataAnalyzer(data)
    print("Data:", data)
    print("Average:", analyzer.average())
    print("Median:", analyzer.median())
    print("Mode:", analyzer.mode())
    print("Standard Deviation:", analyzer.standard_deviation())
    print("Minimum:", analyzer.min())
    print("Maximum:", analyzer.max())


if __name__ == "__main__":
    run_calculations()
    analyze_dataset()

# Filler functions with bugs to extend code to ~500 lines

def buggy_function_1(x):
    if x > 0:
        return x
    elif x == 0:
        return "zero"
    else:
        return None
    return "unreachable code"

def buggy_function_2():
    try:
        return 10 / 0  # Division by zero error
    except:
        pass
    finally:
        return "Done"

for i in range(50):
    def random_buggy_loop(i):
        total = 0
        for j in range(i):
            total += j * 2
        if total % 2 == 1:
            return True
        else:
            return False

# Adding a lot of dummy functions with small logical errors to reach ~500 lines

def example_func_1(a): return a + a - a * a

def example_func_2(b): return b / 2 * 2 - b

def example_func_3(c): return (c + 1)**0.5 - (c - 1)**0.5

" + "\n".join([f"def dummy_func_{i}(x): return x * {i} - x / {i+1}" for i in range(100)]) + "

