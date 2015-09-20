var x = 7;
var z = [1,2,3,4,5,6,7];

function linearSearch(target, array) {
  for (i=0; i<target; i++) {
    if (target === array[i]) {
      console.log(i);
    }
  }
}

linearSearch(x,z);
