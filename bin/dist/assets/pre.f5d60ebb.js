import{b as h,_ as f}from"./index.682428fc.js";/* empty css              *//* empty css               */import{d as g,e as b,bd as v,C as x,aI as t,aG as a,aH as F,be as T,B as w,aM as N,aN as s,F as r,bf as B,bg as C,bh as D}from"./arco.572a1b40.js";import{h as M}from"./vue.adf857d5.js";import"./chart.f83c6245.js";const H={class:"container"},S=["innerHTML"],V=g({__name:"pre",setup(k){const c=M(),o=b(),l={id:Number(c.params.id)};return async function(){try{const n=await h.post("/admin/news",l);o.value=n.data.news}catch(n){throw console.log(n),n}}(),(n,I)=>{const u=F("Breadcrumb"),i=B,_=C,m=v,d=D,y=T;return w(),x("div",H,[t(u,{items:["menu.mSimu","menu.mSimu.pre"]},null,8,["items"]),t(y,null,{default:a(()=>[t(d,null,{default:a(()=>[t(m,{style:{marginTop:"-40px"}},{default:a(()=>[t(i,null,{default:a(()=>{var e;return[N(s((e=o.value)==null?void 0:e.title),1)]}),_:1}),t(_,null,{default:a(()=>{var e,p;return[r("p",null,"\u4F5C\u8005\uFF1A"+s((e=o.value)==null?void 0:e.author),1),r("p",null,"\u65F6\u95F4\uFF1A"+s((p=o.value)==null?void 0:p.date),1)]}),_:1}),t(_,null,{default:a(()=>{var e;return[r("p",{innerHTML:(e=o.value)==null?void 0:e.content},null,8,S)]}),_:1})]),_:1})]),_:1})]),_:1})])}}});const j=f(V,[["__scopeId","data-v-420544db"]]);export{j as default};
