package Problems.Problem_1195_Fizz_Buzz_Multithreaded.java;

import java.util.function.IntConsumer;
import java.util.stream.IntStream;

class FizzBuzz {
    private int n;
    private int current = 1;

    public FizzBuzz(int n) {
        this.n = n;
    }

    // printFizz.run() outputs "fizz".
    public synchronized void fizz(Runnable printFizz) throws InterruptedException {
        while (current <= n) {
            if (current % 3 == 0 && current % 5 != 0) {
                printFizz.run();
                current++;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public synchronized void buzz(Runnable printBuzz) throws InterruptedException {
        while (current <= n) {
            if (current % 5 == 0 && current % 3 != 0) {
                printBuzz.run();
                current++;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public synchronized void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        while (current <= n) {
            if (current % 3 == 0 && current % 5 == 0) {
                printFizzBuzz.run();
                current++;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public synchronized void number(IntConsumer printNumber) throws InterruptedException {
        while (current <= n) {
            if (current % 3 != 0 && current % 5 != 0) {
                printNumber.accept(current);
                current++;
                notifyAll();
            } else {
                wait();
            }
        }
    }

    public static void main(String[] args) {
        Runnable printFizz = new Runnable() {
            public void run() {
                System.out.println("fizz");
            }
        };
        Runnable printBuzz = new Runnable() {
            public void run() {
                System.out.println("buzz");
            }
        };
        Runnable printFizzBuzz = new Runnable() {
            public void run() {
                System.out.println("fizzbuzz");
            }
        };
        IntConsumer printNumber = new IntConsumer() {
            public void accept(int x) {
                System.out.println(x);
            }
        };

        FizzBuzz sol = new FizzBuzz(5);

        // Threads for each operation
        Thread fizzThread = new Thread(() -> {
            try {
                sol.fizz(printFizz);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });
        Thread buzzThread = new Thread(() -> {
            try {
                sol.buzz(printBuzz);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });
        Thread fizzBuzzThread = new Thread(() -> {
            try {
                sol.fizzbuzz(printFizzBuzz);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });
        Thread numberThread = new Thread(() -> {
            try {
                sol.number(printNumber);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });

        // Start the threads
        fizzThread.start();
        buzzThread.start();
        fizzBuzzThread.start();
        numberThread.start();
    }
}
