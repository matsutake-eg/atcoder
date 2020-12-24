#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::marker::*;
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;

// CAUTION: fastout has bug that it is not work in 'match arm', so, in that case, you should remove it or change logic.
#[proconio::fastout]
fn main() {
    proconio::input! {
        n:usize,
    }

    let mut n = n;
    while n % 2 == 0 {
        n /= 2;
    }

    let mut ans = 0;
    let s = n.sqrt();
    for i in 1..=s {
        if n % i == 0 {
            ans += 2;
        }
    }
    if s * s == n {
        ans -= 1;
    }
    println!("{}", ans * 2);
}
