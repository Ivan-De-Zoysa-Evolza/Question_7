# Loan Calculator (Go)

This is a simple command-line loan calculator written in Go.

## What It Does

It calculates the **monthly payment** based on:

- Loan principal amount
- Annual interest rate (%)
- Loan term in months

## Formula Used
M = P × r × (1 + r)^n / ((1 + r)^n - 1)


Where:
- `M` = Monthly payment
- `P` = Principal
- `r` = Monthly interest rate (annual ÷ 12 ÷ 100)
- `n` = Total number of months

If interest rate is 0, it simplifies to `M = P / n`.

## Rules

- Input must be a valid number with up to 2 decimal places.
- Invalid inputs show an error and ask again.
- User can choose to calculate again after each result.

## Run the Program

```bash
go run main.go


