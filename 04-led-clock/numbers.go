package main

func getNum(num int) ClockNum {

    switch num {
        default: return get0()
        case 1: return get1()
        case 2: return get2()
        case 3: return get3()
        case 4: return get4()
        case 5: return get5()
        case 6: return get6()
        case 7: return get7()
        case 8: return get8()
        case 9: return get9()
    }
}

func get0() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,false,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
    }
}

func get1() ClockNum {

    return ClockNum {
        [3]bool{true,true,false,},
        [3]bool{false,true,false,},
        [3]bool{false,true,false,},
        [3]bool{false,true,false,},
        [3]bool{true,true,true,},
    }
}

func get2() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
        [3]bool{true,false,false,},
        [3]bool{true,true,true,},
    }
}

func get3() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
    }
}

func get4() ClockNum {

    return ClockNum {
        [3]bool{true,false,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
    }
}

func get5() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,false,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
    }
}

func get6() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,false,},
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
    }
}

func get7() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
        [3]bool{false,false,true,},
    }
}

func get8() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
    }
}

func get9() ClockNum {

    return ClockNum {
        [3]bool{true,true,true,},
        [3]bool{true,false,true,},
        [3]bool{true,true,true,},
        [3]bool{false,false,true,},
        [3]bool{true,true,true,},
    }
}

