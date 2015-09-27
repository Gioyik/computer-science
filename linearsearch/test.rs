fn main() {
  let x = 5;
  let z: [i32; 7] = [1,2,3,4,5,6,7];
  let mut i = 0;

  loop {
    if z[i] == x {
      println!("{}",i);
      break;
    } else {
      i = i + 1;
    }
  }
}
