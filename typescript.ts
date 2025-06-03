function isHappy(n: number): boolean {
  let sumOfDigits = n
  let iterations = 0

  while (sumOfDigits != 1) {
    sumOfDigits = sumOfDigits
      .toString()
      .split('')
      .map(element => Number(element) ** 2)
      .reduce((response, current) => response + current)

      if (iterations > 20) return false

      iterations++
  }
  return true
}
