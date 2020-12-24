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
        a:[usize;n],
    }

    let a = a.into_iter().filter(|&v| v != x).collect::<Vec<usize>>();
    let l = a.len();
    for (i, v) in a.into_iter().enumerate() {
        if i < l {
            print!("{} ", v);
        } else {
            println!("{}", v);
        }
    }
}
