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
        a:usize,
        b:usize,
        c:usize,
        d:usize,
    }

    let s = a + b + c + d;
    for v in &[a, b, c, d, a + b, a + c, a + d, b + c, b + d, c + d] {
        if *v == s - v {
            println!("{}", "Yes");
            return;
        }
    }
    println!("{}", "No");
}
