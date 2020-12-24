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
        x:usize,
        vp: [(usize,usize);n],
    }

    let mut sum = 0;
    for (i, (v, p)) in vp.into_iter().enumerate() {
        sum += v * p;
        if sum > x * 100 {
            println!("{}", i + 1);
            return;
        }
    }
    println!("{}", -1);
}
