package utl

import "time"

// MapExpire maps an enum value to a formatted date time
// Value can be one of
// 0 -> None
// 1 -> 1h
// 2 -> 12h
// 3 -> 1d
// 4 -> 7d
func MapExpire(value int) string {
    t := time.Now()
    var timeMap = map[int]time.Time{
        0: time.UnixMilli(0000000000), // Hardcoded epoch 0
        1: t.Add(time.Hour),           // 1h
        2: t.Add(time.Hour * 12),      // 12h
        3: t.Add(time.Hour * 24),      // 1d (24h)
        4: t.Add(time.Hour * 24 * 7),  // 7d (24h * 7)
    }
    t = timeMap[value]

    return t.Format("2006-01-02T15:04:05")
}
