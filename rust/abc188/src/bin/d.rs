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
        c_p:i64,
        abc: [(i64,i64,i64);n],
    }

    let mut ev = Vec::new();
    for (a, b, c) in abc {
        ev.push((a - 1, c));
        ev.push((b, -c));
    }
    ev.sort_by(|a, b| (&a.0).cmp(&b.0));

    let mut ans = 0;
    let mut sum = 0;
    let mut t = 0;
    for (x, y) in ev {
        if x != t {
            ans += min(c_p, sum) * (x - t);
            t = x;
        }
        sum += y;
    }
    println!("{}", ans);
}
