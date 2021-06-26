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
        m:usize,
        x:Usize1,
        y:Usize1,
        abtk: [(Usize1,Usize1,usize,usize);m],
    }

    let mut g = vec![vec![]; n];
    for (a, b, t, k) in abtk {
        g[a].push((b, t, k));
        g[b].push((a, t, k));
    }

    const INF: usize = std::usize::MAX;
    let mut ds = vec![INF; n];
    ds[x] = 0;
    let mut pq = BinaryHeap::new();
    pq.push((Reverse(0), x));
    while let Some((Reverse(d), v)) = pq.pop() {
        if d > ds[v] {
            continue;
        }
        for &(nx, t, k) in g[v].iter() {
            let d = (d + k - 1) / k * k + t;
            if d < ds[nx] {
                ds[nx] = d;
                pq.push((Reverse(ds[nx]), nx));
            }
        }
    }
    if ds[y] == INF {
        println!("-1");
    } else {
        println!("{}", ds[y]);
    }
}
