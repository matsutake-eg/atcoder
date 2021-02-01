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
        s:Chars,
    }

    if n % 2 == 1 {
        println!("{}", -1);
        return;
    }

    let mut ans = 0;
    let m = n / 2;
    for i in 0..m {
        if s[i] != s[m + i] {
            ans += 1;
        }
    }
    println!("{}", ans);
}
