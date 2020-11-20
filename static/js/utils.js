function sec2DayHourMinSec(value) {
    let theTime = parseInt(value);// 秒
    let minute = 0;// 分
    let hour = 0;// 小时
    let day = 0;

    if (theTime > 60) {
        minute = parseInt(theTime / 60);
        theTime = parseInt(theTime % 60);
        if (minute > 60) {
            hour = parseInt(minute / 60);
            minute = parseInt(minute % 60);
            if (hour > 24) {
                day = parseInt(hour / 24);
                hour = parseInt(hour % 24);
            }
        }
    }
    let result = "" + parseInt(theTime) + "秒";
    result = "" + parseInt(minute) + "分" + result;

    result = "" + parseInt(hour) + "时" + result;

    result = "" + parseInt(day) + "天" + result;

    return result;
}

function timestamp2datetime(value) {
    let result = new Date(value);
    let year = result.getFullYear()
    let month = result.getMonth()
    let day = result.getDay()
    let hour = result.getHours()
    let minute = result.getMinutes()
    let second = result.getSeconds()
    let resultStr = year + "-" + formatNumber(month) + "-" + formatNumber(day) + " " + formatNumber(hour) + ":" + formatNumber(minute) + ":" + formatNumber(second)
    return resultStr
}

function formatNumber(number) {
    if (number < 10) {
        return "0" + number
    } else {
        return number
    }
}