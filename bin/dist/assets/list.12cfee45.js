import{b as v,_ as N}from"./index.682428fc.js";/* empty css               *//* empty css              *//* empty css               *//* empty css              *//* empty css               *//* empty css               */import{d as k,e as m,aH as I,B as x,C as $,F as c,aI as o,aG as n,aM as s,aW as y,bn as L,b9 as S,ba as V}from"./arco.572a1b40.js";import{f as D}from"./vue.adf857d5.js";import"./chart.f83c6245.js";const b=r=>(S("data-v-dc842dba"),r=r(),V(),r),M={class:"container"},O={class:"list"},P=b(()=>c("h4",null,"banner\u65B0\u95FB\u5217\u8868",-1)),T=b(()=>c("h4",null,"\u5176\u4ED6\u65B0\u95FB",-1)),q=k({__name:"list",setup(r){const E=D(),_=(e,t)=>{E.push({name:e,params:{id:t}})},f={pageSize:5},g=[{title:"id",dataIndex:"id"},{title:"\u9898\u76EE",dataIndex:"title"},{title:"\u65F6\u95F4",dataIndex:"date"},{title:"\u4F5C\u8005",dataIndex:"author"},{title:"\u9884\u89C8",slotName:"pre",width:80},{title:"\u4FEE\u6539",slotName:"edit",width:80},{title:"\u5220\u9664",slotName:"delete",width:80}],p=m([{id:1,title:"banner\u65B0\u95FB1",date:"20230721",author:"\u674E\u5148\u751F"}]),B=[{title:"id",dataIndex:"id"},{title:"\u9898\u76EE",dataIndex:"title"},{title:"\u65F6\u95F4",dataIndex:"date"},{title:"\u4F5C\u8005",dataIndex:"author"},{title:"\u9884\u89C8",slotName:"pre",width:80},{title:"\u4FEE\u6539",slotName:"edit",width:80},{title:"\u5220\u9664",slotName:"delete",width:80}],h=m([{id:1,title:"banner\u65B0\u95FB1",date:"20230721",author:"\u674E\u5148\u751F"}]),i=m([]);(async function(){try{const e=await v.post("/admin/news/banner/all");if(i.value=e.data.news_list,i.value.length){p.value=[];for(let t=0;t<i.value.length;t+=1)p.value.push({id:i.value[t].news_id,title:i.value[t].title,date:i.value[t].date,author:i.value[t].author})}}catch(e){throw console.log(e),e}})();const u=m([]);(async function(){try{const e=await v.post("/admin/news/all");if(u.value=e.data.news_list,u.value.length){h.value=[];for(let t=0;t<u.value.length;t+=1)h.value.push({id:u.value[t].news_id,title:u.value[t].title,date:u.value[t].date,author:u.value[t].author})}}catch(e){throw console.log(e),e}})();const w=async function(e){const t={id:e};await v.post("/admin/news/del",t);debugger};return(e,t)=>{const F=I("Breadcrumb"),l=y,C=L;return x(),$("div",M,[c("div",null,[o(F,{items:["menu.mNotice","menu.mNotice.list"]},null,8,["items"])]),c("div",O,[P,o(C,{columns:g,data:p.value,pagination:f},{pre:n(({record:a})=>[o(l,{onClick:d=>_("NoticePre",a.id)},{default:n(()=>[s("\u9884\u89C8")]),_:2},1032,["onClick"])]),edit:n(({record:a})=>[o(l,{onClick:d=>_("NoticeEdit",a.id)},{default:n(()=>[s("\u4FEE\u6539")]),_:2},1032,["onClick"])]),delete:n(({record:a})=>[o(l,{onClick:d=>e.$modal.info({title:"\u5220\u9664",content:a.id,onOk:()=>w(a.id)})},{default:n(()=>[s("\u5220\u9664")]),_:2},1032,["onClick"])]),_:1},8,["data"])]),c("div",null,[T,o(C,{columns:B,data:h.value,pagination:f},{pre:n(({record:a})=>[o(l,{onClick:d=>_("NoticePre",a.id)},{default:n(()=>[s("\u9884\u89C8")]),_:2},1032,["onClick"])]),edit:n(({record:a})=>[o(l,{onClick:d=>_("NoticeEdit",a.id)},{default:n(()=>[s("\u4FEE\u6539")]),_:2},1032,["onClick"])]),delete:n(({record:a})=>[o(l,{onClick:d=>e.$modal.info({title:"\u5220\u9664",content:a.id,onOk:()=>w(a.id)})},{default:n(()=>[s("\u5220\u9664")]),_:2},1032,["onClick"])]),_:1},8,["data"])])])}}});const Y=N(q,[["__scopeId","data-v-dc842dba"]]);export{Y as default};
