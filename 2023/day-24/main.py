# 327367788702047, 294752664325632, 162080199489440 @ 20, 51, 36
# 349323332347395, 429135322811787, 397812423558610 @ -96, -480, -782
# 342928768632768, 275572250031810, 310926883862869 @ -69, 104, -510
# 308103110384633, 244954649487980, 207561383118617 @ 220, 463, -401

from sympy import symbols, Eq, solve

# Define the symbols
pxs, vxs, pys, vys, pzs, vzs, a1, a2, a3 = symbols('pxs vxs pys vys pzs vzs a1 a2 a3')

# Define the equations
eq1 = Eq(pxs + vxs * a1, 327367788702047 + 20 * a1)
eq2 = Eq(pys + vys * a1, 294752664325632 + 51 * a1)
eq3 = Eq(pzs + vzs * a1, 162080199489440 + 36 * a1)
eq4 = Eq(pxs + vxs * a2, 349323332347395 - 96 * a2)
eq5 = Eq(pys + vys * a2, 429135322811787 - 480 * a2)
eq6 = Eq(pzs + vzs * a2, 397812423558610 - 782 * a2)
eq7 = Eq(pxs + vxs * a3, 342928768632768 - 69 * a3)
eq8 = Eq(pys + vys * a3, 275572250031810 + 104 * a3)
eq9 = Eq(pzs + vzs * a3, 310926883862869 - 510 * a3)

# eq1 = Eq(pxs + vxs * a1, 19 - 2 * a1)
# eq2 = Eq(pys + vys * a1, 13 + a1)
# eq3 = Eq(pzs + vzs * a1, 30 - 2 * a1)
# eq4 = Eq(pxs + vxs * a2, 18 - a2)
# eq5 = Eq(pys + vys * a2, 19 - a2)
# eq6 = Eq(pzs + vzs * a2, 22 - 2 * a2)
# eq7 = Eq(pxs + vxs * a3, 20 - 2 * a3)
# eq8 = Eq(pys + vys * a3, 25 - 2 * a3)
# eq9 = Eq(pzs + vzs * a3, 34 - 4 * a3)


# Solve the system of equations
solution = solve((eq1, eq2, eq3, eq4, eq5, eq6, eq7, eq8, eq9), (pxs, vxs, pys, vys, pzs, vzs, a1, a2, a3))

print(solution)

print(solution[0][0] + solution[0][2] + solution[0][4])