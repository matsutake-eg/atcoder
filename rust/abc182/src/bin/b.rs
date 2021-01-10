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
        n:usize,
        a:[usize;n],
    }

    let mut ans = 0;
    let mut g = 0;
    for i in 2..=1000 {
        let mut t = 0;
        for v in &a {
            if i > *v {
                continue;
            }
            if v % i == 0 {
                t += 1;
            }
        }
        if t > g {
            g = t;
            ans = i;
        }
    }
    println!("{}", ans);
}
