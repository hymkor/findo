readnum - Read float from bufio.Reader
======================================

    func readnum.Float(r *bufio.Reader) (float64, bool, error)

    '' -> 0.000000(false)EOF
    '.' -> 0.000000(true)EOF
    '.1' -> 0.100000(true)EOF
    '3' -> 3.000000(true)EOF
    '3 ' -> 3.000000(true)EOF
    '-3' -> -3.000000(true)EOF
    '-3 ' -> -3.000000(true)EOF
    '3.14' -> 3.140000(true)EOF
    '3.14 ' -> 3.140000(true)<nil>
    '3.14e2' -> 314.000000(true)EOF
    '3.14e2 ' -> 314.000000(true)<nil>
    '3.14e-2' -> 0.031400(true)EOF
    '3.14e-2 ' -> 0.031400(true)<nil>
    '-3.14e+1' -> -31.400000(true)EOF
    '-3.14e+1 ' -> -31.400000(true)<nil>
    '-0.314e+1' -> -3.140000(true)EOF
    '-0.314e+1 ' -> -3.140000(true)<nil>
