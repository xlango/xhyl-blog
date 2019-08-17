<template>
<div class="blog-list">
   <HeadNav style="margin-bottom:120px;"></HeadNav>
   <div class="container-fluid">
        <div class="row">
            <div class="col-md-1"></div>
            <div class="col-md-7">
                <div class="container-fluid">
                    <div class="row">
                        <Card class="card container">
                            <div class="row">
                                <div class="col-md-8">
                                    <img src="@/assets/avatar.png" style="width:120px;height:120px;margin:5px;border-radius:10%;margin-top:-60px;"/>
                                    <span style="margin:5px;font-size: 20px;"><b>张三</b>
                                        <img src="https://file.iviewui.com/dev-dist/3eb78a60b677c55480495de82d69c8bc.png" style="width:20px;height:20px;margin-left:5px;">
                                        <span style="font-size: 10px;">普通会员</span>
                                    </span>
                                </div>
                            </div>
                            <div class="row">
                                <Button style="float:right;" size="small">编辑个人资料</Button>
                            </div>
                        </Card>
                        <Card class="card">
                            <Tabs>
                                <TabPane label="文章 0" icon="md-list-box">
                                    <div align="center">
                                        <img style="width:420px;height:250px;" src="https://file.iviewui.com/dev-dist/a8899afd8b36404e2829f9aaa18881b8.png">
                                        <div style="text-align: center;color: #b2b2b2;">暂时没有内容</div>
                                    </div>
                                </TabPane>
                                <TabPane label="提问 0" icon="md-help-circle">
                                    <div align="center">
                                        <img style="width:420px;height:250px;" src="https://file.iviewui.com/dev-dist/a8899afd8b36404e2829f9aaa18881b8.png">
                                        <div style="text-align: center;color: #b2b2b2;">暂时没有内容</div>
                                    </div>
                                </TabPane>
                                <TabPane label="收藏 0" icon="ios-heart">
                                    <div align="center">
                                        <img style="width:420px;height:250px;" src="https://file.iviewui.com/dev-dist/a8899afd8b36404e2829f9aaa18881b8.png">
                                        <div style="text-align: center;color: #b2b2b2;">暂时没有内容</div>
                                    </div>
                                </TabPane>
                            </Tabs>
                        </Card>
                        <Button class="btnmore" title="点击加载更多"><Icon type="md-arrow-down" /></Button>
                    </div>
                </div>
                
            </div>
            <div class="col-md-3">
                <Card class="card-slider">
                    <Input suffix="ios-search" placeholder="搜索" style="width: 100%" />
                </Card>
                <Card class="card-slider">
                    <DatePicker :value="nowTime" format="yyyy年MM月dd日" type="date" placeholder="Select date" style="width: 100%;"></DatePicker>
                </Card>
                <Card class="card-slider">
                    <Row>
                        <Icon type="ios-options" /><span style="margin:5px;">发布历程：</span>
                    </Row>
                    <Row>
                        <hr  style="height:1px;border:none;border-top:1px solid rgb(136, 135, 135);margin:10px;"/>
                    </Row>
                    <Row>
                        <Timeline>
                            <TimelineItem color="green">发布1.0版本</TimelineItem>
                            <TimelineItem color="green">发布2.0版本</TimelineItem>
                            <TimelineItem color="red">严重故障</TimelineItem>
                            <TimelineItem color="blue">发布3.0版本</TimelineItem>
                        </Timeline>
                    </Row>
                </Card>
                
                <Button type="success"  class="card-slider" long><Icon type="ios-create-outline"/>写文章</Button>
            </div>
            <div class="col-md-1"></div>
        </div>
    </div>
</div>
</template>
<script>
import BlogListCard from "@/components/BlogListCard";
import HeadNav from '@/components/HeadNav'
export default {
  name: "blog-list",
  components:{
        HeadNav
  },
  data () {
            return {
                nowTime:new Date().toLocaleString(),
            }
        },
		methods:{

		},
		mounted(){

		}
};
</script>
<style scoped>
.blog-list {
  height: 100%;
}
.card {
  width: 90%;
  margin-top: 20px;
  margin-left: 90px;
}
.card-slider {
  width: 80%;
  margin-top: 20px;
}

.btntype {
  height: 23px;
  width: 65px;
  margin: 5px;
  display: inline-block;
  font-size: 14px;
  font-weight: 400;
  text-align: center;
  background-image: none;
  color: #333;
  background-color: white;
    color: black;
    border: 1px solid #e7e7e7;
}
.btnmore {
  width: 90%;
  margin-top: 20px;
  margin-left: 90px;
}

</style>