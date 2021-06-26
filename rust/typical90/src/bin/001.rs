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

    let mut b = vec![0; n + 1];
    b[0] = a[0];
    for i in 1..n {
        b[i] = a[i] - a[i - 1];
    }
    b[n] = l - a[n - 1];

    let mut ans = l / 2;
    let mut total = 0;
    let mut cnt = 0;
    loop {
        for v in &b {
            total += v;
            if total >= ans {
                cnt += 1;
                total = 0;
            }
        }
        if cnt < k {
            ans /= 2;
        } else if cnt > k {
            ans = (l - ans) / 2;
        } else {
            break;
        }
    }
    println!("{}", ans);
}
