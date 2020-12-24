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
        m:usize,
        ab: [(Usize1,Usize1);m],
        k:usize,
        cd: [(Usize1,Usize1);k],
    }

    let mut ans = 0;
    for i in 0..(1 << k) {
        let mut ext = vec![false; n];
        for j in 0..k {
            if i >> j & 1 == 1 {
                ext[cd[j].0] = true;
            } else {
                ext[cd[j].1] = true;
            }
        }

        let mut sum = 0;
        for &(a, b) in &ab {
            if ext[a] && ext[b] {
                sum += 1;
            }
        }
        ans = max(ans, sum);
    }
    println!("{}", ans);
}
