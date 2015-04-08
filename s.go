package s	// Subroutines

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "runtime/debug"
)

type ChannelBuf struct {
    Str     string
    channel chan string
    eof     bool
}

func (this *ChannelBuf) Next() bool {
    if !this.eof {
        if line, ok := <-this.channel; ok {
            this.Str = line
            return true
        } else {
            this.eof = true
        }
    }
    return false
}

func Open(fname string) *ChannelBuf {
    ch := make(chan string, 32) // Some arbitrary amount of readahead

    file, err := os.Open(fname)
    if err != nil {
        debug.PrintStack()
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)

    go func() {
        defer func() {
            file.Close()
            close(ch)
        }()
        for scanner.Scan() {
            ch <- scanner.Text()
        }
        if err := scanner.Err(); err != nil {
            debug.PrintStack()
            log.Fatal(err)
        }
    }()

    return &ChannelBuf{"", ch, false}
}

func Match(target, re string) []string {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    return match
}

func Match1(target, re string) (bool, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 2 {
        return true, match[1]
    }
    return false, ""
}

func Match2(target, re string) (bool, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 3 {
        return true, match[1], match[2]
    }
    return false, "", ""
}

func Match3(target, re string) (bool, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 4 {
        return true, match[1], match[2], match[3]
    }
    return false, "", "", ""
}

func Match4(target, re string) (bool, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 5 {
        return true, match[1], match[2], match[3], match[4]
    }
    return false, "", "", "", ""
}

func Match5(target, re string) (bool, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 6 {
        return true, match[1], match[2], match[3], match[4], match[5]
    }
    return false, "", "", "", "", ""
}

func Match6(target, re string) (bool, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 7 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6]
    }
    return false, "", "", "", "", "", ""
}

func Match7(target, re string) (bool, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 8 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7]
    }
    return false, "", "", "", "", "", "", ""
}

func Match8(target, re string) (bool, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 9 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8]
    }
    return false, "", "", "", "", "", "", "", ""
}

func Match9(target, re string) (bool, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 10 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9]
    }
    return false, "", "", "", "", "", "", "", "", ""
}

func Match10(target, re string) (bool, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 11 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10]
    }
    return false, "", "", "", "", "", "", "", "", "", ""
}

func Match11(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 12 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11]
    }
    return false, "", "", "", "", "", "", "", "", "", "", ""
}

func Match12(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 13 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match13(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 14 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match14(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 15 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match15(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 16 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14], match[15]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match16(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 17 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14], match[15], match[16]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match17(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 18 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14], match[15], match[16], match[17]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match18(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 19 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14], match[15], match[16], match[17], match[18]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match19(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 20 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14], match[15], match[16], match[17], match[18], match[19]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

func Match20(target, re string) (bool, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string, string) {
    r, error := regexp.Compile(re)
    if error != nil {
        debug.PrintStack()
        log.Fatal(error)
    }
    match := r.FindStringSubmatch(target)
    if len(match) == 21 {
        return true, match[1], match[2], match[3], match[4], match[5], match[6], match[7], match[8], match[9], match[10], match[11], match[12], match[13], match[14], match[15], match[16], match[17], match[18], match[19], match[20]
    }
    return false, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""
}

// vim:ts=4:sw=4:et
