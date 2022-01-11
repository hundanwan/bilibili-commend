var currentIndex = 0
var maxLength = 0
var timer = null
var left = 0

window.onload = function () {
    maxLength = document.getElementById("container").getElementsByTagName("li").length - 1
}
var move = (change) => {
    var lb = document.getElementById("b-left");
    var rb = document.getElementById("b-right");
    lb.onclick = null
    rb.onclick = null
    //console.log(lb, rb)
    currentIndex += change
    if (currentIndex == 4) {
        left = 0
        currentIndex = 1
        container.style.left = 0
        //console.log("reset")
    }
    if (currentIndex == -1) {
        left = -300
        currentIndex = 2
        container.style.left = "-300%"
        //console.log("reset")
    }
    if (change == 1) {
        next()
    } else {
        last()
    }

    console.log(currentIndex)

}
var next = () => {
    var lb = document.getElementById("b-left");
    var rb = document.getElementById("b-right");
    var container = document.getElementById("container")
    clearInterval(timer)

    timer = setInterval(function () {
        //通过百分比来移动left
        if (left > currentIndex * -100) {
            left -= 1
            container.style.left = left + "%"
            console.log(left)
        } else {
            lb.onclick = function () { move(-1) }
            rb.onclick = function () { move(1) }
            // console.log(left, currentIndex * -100)
        }

    }, 1)


}
var last = () => {
    var lb = document.getElementById("b-left");
    var rb = document.getElementById("b-right");
    clearInterval(timer)
    timer = setInterval(function () {
        if (left < currentIndex * -100) {
            left += 1
            container.style.left = left + "%"
        } else {
            lb.onclick = function () { move(-1) }
            rb.onclick = function () { move(1) }
            // console.log(left, currentIndex * -100)
        }
    }, 1)
}

var mobileOpen = () => {
    var navMobile = document.getElementById("nav-mobile");
    navMobile.style.left = 0;
}

var mobileClose = () => {
    var navMobile = document.getElementById("nav-mobile");
    navMobile.style.left = "-50%";
}
