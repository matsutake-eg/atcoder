#![allow(unused_imports)]
use itertools::Itertools;
use itertools_num::ItertoolsNum as _;
use num_integer::*;
use petgraph::*;
use proconio::{fastout, input, marker::*};
use std::cmp::*;
use std::collections::*;
use std::f64::consts::*;
use superslice::Ext as _;

// CAUTION: don't print in 'match arm' because of fastout's bug
#[fastout]
fn main() {
    input! {
        n:usize,
        l:usize,
        k:usize,
        a:[usize;n],
    }

    let (mut left, mut right) = (0, l);
    while right - left > 1 {
        let m = (left + right) / 2;
        let mut pre = 0;
        let mut cnt = 0;
        for v in &a {
            if v - pre >= m && l - v >= m {
                cnt += 1;
                pre = *v;
            }
        }
        if cnt < k {
            right = m;
        } else {
            left = m;
        }
    }
    println!("{}", left);
}
