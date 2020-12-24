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
        p:[usize;n],
    }

    let mut a = vec![false; 200_001];
    let mut ans = 0;
    for v in p {
        a[v] = true;
        while a[ans] {
            ans += 1;
        }
        println!("{}", ans);
    }
}
