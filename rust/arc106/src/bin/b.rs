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
        a:[i64;n],
        b:[i64;n],
        cd: [(Usize1,Usize1);m],
    }

    let mut uft = unionfind::UnionFind::new(n);
    for (c, d) in cd {
        uft.union(c, d);
    }

    let mut cnt = vec![0i64; n];
    for i in 0..n {
        let f = uft.find(i);
        cnt[f] += a[i];
        cnt[f] -= b[i];
    }
    println!(
        "{}",
        if cnt.into_iter().all(|x| x == 0) {
            "Yes"
        } else {
            "No"
        }
    );
}
