

# Exercise 1: Extend the `calculator` Package

-----

## Goal

Learn how to define a function in an external **package** and call that function from the **`main`** function.

-----

## Task

1.  **Extend `calculator/calculator.go`:**

      * Add a new, **exportable** function named **`Multiply`**.
      * The function must accept two `int` values and return their product as an `int`.
      * **Reminder:** Function names must start with a **capital letter** to be exportable\!

2.  **Use the Function in `main.go`:**

      * In the `func main()`, call the function **`calculator.Multiply(4, 5)`**.
      * Store the result in a variable.
      * Print the result along with a description (e.g., "Product:") to the console.

-----

##  Expected Output

In addition to the existing output, the console should display:

```
Product: 20
```