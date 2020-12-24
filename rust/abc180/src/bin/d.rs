#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::fastout;
use proconio::input;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

#[allow(unused)]
const MOD: usize = 1_000_000_007;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        x:usize,
        y:usize,
        a:usize,
        b:usize,
    }

    let mut ans = 0;
    let mut x = x;
    while (a - 1) * x <= b && x * a < y {
        x *= a;
        ans += 1;
    }
    ans += (y - x - 1) / b;
    println!("{}", ans);
}
