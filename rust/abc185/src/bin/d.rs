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
        m:usize,
        a:[usize;m],
    }

    if m == 0 {
        println!(1);
        return;
    }

    let mut a = a;
    a.sort();

    let mut d = Vec::new();
    for i in 0..(m - 1) {
        let v = a[i + 1] - a[i] - 1;
        if v >= 1 {
            d.push(v);
        }
    }
    if d.len() == 0 {
        println!(0);
        return;
    }

    let mut k = *d.iter().min().unwrap();
    if a[0] > 1 {
        k = min(k, a[0] - 1);
    }
    if a[m - 1] < n {
        k = min(k, n - a[m - 1]);
    }

    let v1 = a[0] - 1;
    let v2 = n - a[m - 1];
    let mut ans = (v1 + k - 1) / k + (v2 + k - 1) / k;
    for v in d {
        ans += (v + k - 1) / k;
    }
    println!("{}", ans);
}
