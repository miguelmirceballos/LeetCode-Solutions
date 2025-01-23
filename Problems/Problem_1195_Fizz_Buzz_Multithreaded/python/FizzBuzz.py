from typing import Callable
from threading import Thread, Condition

class FizzBuzz:
    def __init__(self, n: int):
        self.n = n
        self.counter = 1
        self.condition = Condition()

    # printFizz() outputs "fizz"
    def fizz(self, printFizz: Callable[[], None]) -> None:
        while self.counter <= self.n:
            with self.condition:
                while self.counter <= self.n and (self.counter % 3 != 0 or self.counter % 5 == 0):
                    self.condition.wait()  # Esperar turno
                if self.counter > self.n:
                    break
                printFizz()
                self.counter += 1
                self.condition.notify_all()  # Notificar a otros hilos

    # printBuzz() outputs "buzz"
    def buzz(self, printBuzz: Callable[[], None]) -> None:
        while self.counter <= self.n:
            with self.condition:
                while self.counter <= self.n and (self.counter % 3 == 0 or self.counter % 5 != 0):
                    self.condition.wait()  # Esperar turno
                if self.counter > self.n:
                    break
                printBuzz()
                self.counter += 1
                self.condition.notify_all()  # Notificar a otros hilos

    # printFizzBuzz() outputs "fizzbuzz"
    def fizzbuzz(self, printFizzBuzz: Callable[[], None]) -> None:
        while self.counter <= self.n:
            with self.condition:
                while self.counter <= self.n and (self.counter % 3 != 0 or self.counter % 5 != 0):
                    self.condition.wait()  # Esperar turno
                if self.counter > self.n:
                    break
                printFizzBuzz()
                self.counter += 1
                self.condition.notify_all()  # Notificar a otros hilos

    # printNumber(x) outputs "x", where x is an integer.
    def number(self, printNumber: Callable[[int], None]) -> None:
        while self.counter <= self.n:
            with self.condition:
                while self.counter <= self.n and (self.counter % 3 == 0 or self.counter % 5 == 0):
                    self.condition.wait()  # Esperar turno
                if self.counter > self.n:
                    break
                printNumber(self.counter)
                self.counter += 1
                self.condition.notify_all()  # Notificar a otros hilos


if __name__ == "__main__":
    
    # Funciones para imprimir
    printFizz = lambda: print("fizz")
    printBuzz = lambda: print("buzz")
    printFizzBuzz = lambda: print("fizzbuzz")
    printNumber = lambda x: print(x)

    # Crear instancia de FizzBuzz
    sol = FizzBuzz(15)

    # Crear hilos
    threadA = Thread(target=sol.fizz, args=(printFizz,))
    threadB = Thread(target=sol.buzz, args=(printBuzz,))
    threadC = Thread(target=sol.fizzbuzz, args=(printFizzBuzz,))
    threadD = Thread(target=sol.number, args=(printNumber,))

    # Iniciar hilos
    threadA.start()
    threadB.start()
    threadC.start()
    threadD.start()

    # Esperar a que todos los hilos terminen
    threadA.join()
    threadB.join()
    threadC.join()
    threadD.join()




