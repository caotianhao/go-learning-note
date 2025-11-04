package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * // golang的对应函数，请直接使用对应的函数 // 不要删除对应的Sleep方法，这里模拟了接口调用的时间。  func GetPageViews(date string) int { time.Sleep(100 * time.Millisecond)  t, err := time.Parse("2006-01-02", date) if err != nil { return 0 }  start, _ := time.Parse("2006-01-02", "2025-10-01") end, _ := time.Parse("2006-01-02", "2025-10-12")  if !t.Before(start) && !t.After(end) { return 0 }  dateNum := dateToNumber(date) views := ((dateNum * 9973) % 30001) + 20000 return views }  func GetAdClicks(date string) int { time.Sleep(100 * time.Millisecond)  t, err := time.Parse("2006-01-02", date) if err != nil { return 0 }  start, _ := time.Parse("2006-01-02", "2025-10-01") end, _ := time.Parse("2006-01-02", "2025-10-12")  if !t.Before(start) && !t.After(end) { return 0 }  dateNum := dateToNumber(date) clicks := ((dateNum * 7919) % 5001) + 5000 return clicks }  func dateToNumber(date string) int { t, err := time.Parse("2006-01-02", date) if err != nil { return 0 }  year := t.Year() month := int(t.Month()) day := t.Day()  return year*10000 + month*100 + day }  // Java版本的给定函数  private static final DateTimeFormatter DATE_FORMAT = DateTimeFormatter.ofPattern("yyyy-MM-dd");  public static int getPageViews(String date) { try { Thread.sleep(100); } catch (InterruptedException e) { Thread.currentThread().interrupt(); }  LocalDate dt = LocalDate.parse(date, DATE_FORMAT); LocalDate start = LocalDate.parse("2025-10-01", DATE_FORMAT); LocalDate end = LocalDate.parse("2025-10-12", DATE_FORMAT);  if (!dt.isBefore(start) && !dt.isAfter(end)) { return 0; }  long dateNum = dateToNumber(date); return (int)(((dateNum * 9973) % 30001) + 20000); }  public static int getAdClicks(String date) { try { Thread.sleep(100); } catch (InterruptedException e) { Thread.currentThread().interrupt(); }  LocalDate dt = LocalDate.parse(date, DATE_FORMAT); LocalDate start = LocalDate.parse("2025-10-01", DATE_FORMAT); LocalDate end = LocalDate.parse("2025-10-12", DATE_FORMAT);  if (!dt.isBefore(start) && !dt.isAfter(end)) { return 0; }  long dateNum = dateToNumber(date); return (int)(((dateNum * 7919) % 5001) + 5000); }  public static int dateToNumber(String date) { LocalDate dt = LocalDate.parse(date, DATE_FORMAT); int year = dt.getYear(); int month = dt.getMonthValue(); int day = dt.getDayOfMonth(); return year * 10000 + month * 100 + day; }
 * @param startDate string字符串 开始日期（包含）
 * @param endDate string字符串 结束日期（包含）
 * @return string字符串一维数组
 */

func GetPageViews(date string) int {
	time.Sleep(100 * time.Microsecond)
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	start, _ := time.Parse("2006-01-02", "2025-10-01")
	end, _ := time.Parse("2006-01-02", "2025-10-12")
	if !t.Before(start) && !t.After(end) {
		return 0
	}
	dateNum := dateToNumber(date)
	views := ((dateNum * 9973) % 30001) + 20000
	return views
}

func GetAdClicks(date string) int {
	time.Sleep(100 * time.Microsecond)
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	start, _ := time.Parse("2006-01-02", "2025-10-01")
	end, _ := time.Parse("2006-01-02", "2025-10-12")
	if !t.Before(start) && !t.After(end) {
		return 0
	}
	dateNum := dateToNumber(date)
	clicks := ((dateNum * 7919) % 5001) + 5000
	return clicks
}

func dateToNumber(date string) int {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	return year*10000 + month*100 + day
}

func CalculateWeeklyClickRate(startDate string, endDate string) []string {
	divRes := DivDate(startDate, endDate)
	fmt.Println(divRes)
	var res []string
	var wg sync.WaitGroup
	sem := make(chan struct{}, 10)

	for _, v := range divRes {
		s, e := v[0], v[1]
		start, _ := time.Parse("2006-01-02", s)
		end, _ := time.Parse("2006-01-02", e)
		vv, g := 0, 0
		for current := start; !current.After(end); current = current.AddDate(0, 0, 1) {
			wg.Add(1)
			sem <- struct{}{}
			go func(cc time.Time, vvv, gg *int) {
				defer wg.Done()
				defer func() { <-sem }()
				*vvv += GetPageViews(cc.Format("2006-01-02"))
				*gg += GetAdClicks(cc.Format("2006-01-02"))
			}(current, &vv, &g)
		}
		wg.Wait()
		if vv == 0 || g == 0 {
			res = append(res, "0.00%")
		} else {
			res = append(res, fmt.Sprintf("%.2f%%", float64(g)/float64(vv)*100))
		}
	}

	return res
}

func DivDate(startDate string, endDate string) [][2]string {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	var result [][2]string
	current := start

	for !current.After(end) {
		weekday := current.Weekday()
		diff := (int(weekday) - 1 + 7) % 7

		mm := current.AddDate(0, 0, -diff)
		ss := mm.AddDate(0, 0, 6)

		ac := mm

		if ac.Before(start) {
			ac = start
		}
		as := ss

		if as.After(end) {
			as = end
		}

		result = append(result, [2]string{
			ac.Format("2006-01-02"),
			as.Format("2006-01-02"),
		})

		current = ss.AddDate(0, 0, 1)
	}

	return result
}

func main() {
	fmt.Println(CalculateWeeklyClickRate("2025-09-01", "2025-10-12"))
}
