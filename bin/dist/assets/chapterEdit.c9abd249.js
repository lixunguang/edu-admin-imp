import{b as v,_ as q}from"./index.682428fc.js";/* empty css              *//* empty css               *//* empty css               *//* empty css               *//* empty css              *//* empty css               *//* empty css              *//* empty css               */import{d as D,e as f,aW as R,bn as H,bi as L,b2 as O,C as U,aI as t,aG as a,aH as T,a$ as j,B as z,F as b,aM as i,aN as G,bh as W,be as J,bj as K,b9 as P,ba as Q}from"./arco.572a1b40.js";import{f as X,h as Y}from"./vue.adf857d5.js";import"./chart.f83c6245.js";const Z=m=>(P("data-v-aad19ca7"),m=m(),Q(),m),ee={class:"container"},ae=Z(()=>b("p",null,"\u7AE0\u8282\u5217\u8868",-1)),te={style:{marginLeft:"40px",marginRight:"40px"}},se=D({__name:"chapterEdit",setup(m){const E=X(),r=Y(),g=r.params.id,w=(o,e)=>{E.push({name:o,params:{courseId:g,chapterId:e}})},I={pageSize:5},x=[{title:"ID",dataIndex:"id"},{title:"\u6807\u9898",dataIndex:"name"},{title:"\u63CF\u8FF0",dataIndex:"desc"},{title:"\u4FEE\u6539",slotName:"edit",width:80},{title:"\u5220\u9664",slotName:"delete",width:80},{title:"\u7F16\u8F91\u8BE6\u60C5",slotName:"editMore",width:80}],d=f([]),c=f(!1),h=f(!1),s=f({id:1,name:"\u7B2C\u4E00\u8282",desc:"\u7B2C\u4E00\u8282\u63CF\u8FF0"});(async function(){var o;try{const e={id:Number(r.params.id)},n=await v.post("/admin/course/lesson/all",e);for(let l=0;l<((o=n.data)==null?void 0:o.length);l+=1)d.value.push({id:n.data[l].id,name:n.data[l].title,desc:n.data[l].desc})}catch(e){throw console.log(e),e}})();const k=()=>{c.value=!0,h.value=!0},y=(o,e,n)=>{c.value=!0,h.value=!1,s.value.id=o,s.value.desc=n,s.value.name=e},B=async function(o){var n;for(let l=0;l<((n=d.value)==null?void 0:n.length);l+=1)d.value[l].id===o&&d.value.splice(l,1);const e={course_id:Number(r.params.id),lesson_id:Number(o)};await v.post("/admin/course/lesson/del",e)},N=async function(){if(h.value){const o={course_id:Number(r.params.id),lesson_title:s.value.name,lesson_desc:s.value.desc},e=await v.post("/admin/course/lesson/add",o);debugger;const n=e.data.lesson_id;d.value.push({id:n,name:s.value.name,desc:s.value.desc})}else{for(let e=0;e<d.value.length;e+=1)d.value[e].id===s.value.id&&(d.value[e].name=s.value.name,d.value[e].desc=s.value.desc);const o={course_id:Number(r.params.id),lesson_id:s.value.id,title:s.value.name,desc:s.value.desc};await v.post("/admin/course/lesson/update",o)}c.value=!1},A=()=>{c.value=!1};return(o,e)=>{const n=T("Breadcrumb"),l=R,p=W,_=J,M=H,V=K,S=L,C=O,$=j;return z(),U("div",ee,[t(n,{items:["menu.mSimu","menu.mSimu.chapterEdit"]},null,8,["items"]),t(S,null,{default:a(()=>[t(V,null,{default:a(()=>[ae,b("div",te,[t(_,null,{default:a(()=>[t(p,{span:2},{default:a(()=>[t(l,{type:"primary",onClick:e[0]||(e[0]=u=>k())},{default:a(()=>[i(" \u589E\u52A0\u7AE0\u8282 ")]),_:1})]),_:1})]),_:1}),t(M,{columns:x,data:d.value,pagination:I},{pre:a(({record:u})=>[t(l,null,{default:a(()=>[i("\u9884\u89C8"+G(u.id),1)]),_:2},1024)]),edit:a(({record:u})=>[t(l,{onClick:F=>y(u.id,u.name,u.desc)},{default:a(()=>[i("\u4FEE\u6539")]),_:2},1032,["onClick"])]),delete:a(({record:u})=>[t(l,{onClick:F=>o.$modal.info({title:"\u5220\u9664",content:u.id,onOk:()=>B(u.id)})},{default:a(()=>[i("\u5220\u9664")]),_:2},1032,["onClick"])]),editMore:a(({record:u})=>[t(l,{onClick:F=>w("SimuChapterDetailEdit",u.id)},{default:a(()=>[i("\u7F16\u8F91\u8BE6\u60C5")]),_:2},1032,["onClick"])]),_:1},8,["data"])])]),_:1})]),_:1}),t($,{visible:c.value,"onUpdate:visible":e[3]||(e[3]=u=>c.value=u),onOk:N,onCancel:A},{title:a(()=>[i(" \u65B0\u589E/\u7F16\u8F91\u7AE0\u8282\u57FA\u672C\u4FE1\u606F ")]),default:a(()=>[t(_,null,{default:a(()=>[i("\u7AE0\u8282\u57FA\u672C\u4FE1\u606F")]),_:1}),t(_,null,{default:a(()=>[t(p,{span:4},{default:a(()=>[i("\u6807\u9898\uFF1A")]),_:1}),t(p,{span:4},{default:a(()=>[t(C,{modelValue:s.value.name,"onUpdate:modelValue":e[1]||(e[1]=u=>s.value.name=u),style:{width:"320px"},"allow-clear":""},null,8,["modelValue"])]),_:1})]),_:1}),t(_,null,{default:a(()=>[t(p,{span:4},{default:a(()=>[i("\u63CF\u8FF0\uFF1A")]),_:1}),t(p,{span:4},{default:a(()=>[t(C,{modelValue:s.value.desc,"onUpdate:modelValue":e[2]||(e[2]=u=>s.value.desc=u),style:{width:"320px"},"allow-clear":""},null,8,["modelValue"])]),_:1})]),_:1})]),_:1},8,["visible"])])}}});const he=q(se,[["__scopeId","data-v-aad19ca7"]]);export{he as default};
