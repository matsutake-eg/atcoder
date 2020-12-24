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
        a:[usize;1<<n],
    }

    let mn = min(
        a[..(a.len() / 2)].iter().max().unwrap(),
        a[(a.len() / 2)..].iter().max().unwrap(),
    );
    println!("{}", a.iter().position(|&x| x == *mn).unwrap() + 1);
}
