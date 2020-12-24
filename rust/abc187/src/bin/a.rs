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
        a:usize,
        b:usize,
    }

    let mut a_sum = 0;
    let mut a = a;
    while a > 0 {
        a_sum += a % 10;
        a /= 10;
    }
    let mut b_sum = 0;
    let mut b = b;
    while b > 0 {
        b_sum += b % 10;
        b /= 10;
    }
    println!("{}", max(a_sum, b_sum));
}
