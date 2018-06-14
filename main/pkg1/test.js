function fib1(n) {
  var result;
  if (n <=1 ) {
    return 1;
  } else {
    result = fib1(n - 1) + fib1(n - 2)
  }
  return result;
}
console.time("fib1--普通的:")
console.log("fib1--普通的:", fib1(41))
console.timeEnd("fib1--普通的:");



var a = [];
console.time("fib2--内存缓存的:");
console.log('fib2--内存缓存的:', fib2(41));
console.timeEnd("fib2--内存缓存的:");

function fib2(n) {
  var result;
  for (var i = 0; i <= n; i++) {
    i === n ? (result = _fib2(n)) : _fib2(n);
  }
  return result;
}
function _fib2 (n) {
  if (a[n]) {
    return a[n];
  }
  var result;
  if (n <= 1 ) {
    return 1;
  } else {
    result = _fib2(n - 1) + _fib2(n - 2)
  }
  a[n] = result;
  return result;
}