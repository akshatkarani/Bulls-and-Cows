package main

func GetSet(num1 int, num2 int) set {
    guess := num1
    bulls := 0
    cows := 0
    var a, b [4]int
    for i := 0; i < 4; i++ {
        a[3-i] = num1 % 10
        b[3-i] = num2 % 10
        num1 = num1 / 10
        num2 = num2 / 10
    }
    if a[0] == b[0] {
        bulls += 1
    } else if a[0] == b[1] || a[0] == b[2] || a[0] == b[3] {
        cows += 1
    }
    if a[1] == b[1] {
        bulls += 1
    } else if a[1] == b[0] || a[1] == b[2] || a[1] == b[3] {
        cows += 1
    }
    if a[2] == b[2] {
        bulls += 1
    } else if a[2] == b[1] || a[2] == b[0] || a[2] == b[3] {
        cows += 1
    }
    if a[3] == b[3] {
        bulls += 1
    } else if a[3] == b[1] || a[3] == b[2] || a[3] == b[0] {
        cows += 1
    }
    return set{guess, bulls, cows}
}
