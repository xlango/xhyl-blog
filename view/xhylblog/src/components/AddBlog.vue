<template>
<div class="add-blog">
  <HeadNav></HeadNav>
  <div class="container-fluid" style="margin-top: 50px;">
        <div class="row">
            <div class="col-md-1"></div>
            <div class="col-md-7">
                <div class="container-fluid">
                    <div class="row">
                        <Card class="card container">
                            <div class="row">
                                <Form :model="formItem" :label-width="80">
                                    <FormItem label="标题：">
                                        <Input v-model="blog.Article.Title" placeholder="Title" clearable></Input>
                                    </FormItem>
                                    <FormItem label="类型：" style="width: 50%;">
                                        <Select v-model="selectTypeId" placeholder="请选择类型">
                                            <Option v-for="t in types" v-bind:key="t.Id" v-bind:value="t.Id">
                                                {{t.TypeName}}
                                            </Option>
                                        </Select>
                                    </FormItem>
                                    <FormItem label="可见：">
                                        <i-switch v-model="formItem.switch" size="large">
                                            <span slot="open">On</span>
                                            <span slot="close">Off</span>
                                        </i-switch>
                                    </FormItem>
                                    <FormItem label="描述：">
                                        <Input v-model="blog.Article.Content" type="textarea" :autosize="{minRows: 2,maxRows: 5}" placeholder="Enter something..." clearable></Input>
                                    </FormItem>
                                </Form>
                            </div>
                        </Card>
                        <Card class="card container">
                            <Form ref="Paragraphs" :model="Paragraphs" :label-width="80" style="width: 100;">
                              <FormItem
                                      v-for="(item, index) in Paragraphs.items"
                                      v-if="item.status"
                                      :key="index"
                                      :label="'章节 ' + item.index">
                                  <Row>
                                      <Col span="18">
                                          <Input type="text" v-model="item.Title" placeholder="请输入章节标题" clearable></Input>
                                      </Col>
                                      <Col span="4" offset="1">
                                          <Button @click="removeParagh(index)" type="error">删除</Button>
                                      </Col>
                                  </Row>
                                  <Row style="margin-top:10px;">
                                      <Col span="18">
                                          <Input v-model="item.Content" type="textarea" :autosize="{minRows: 5,maxRows: 8}" placeholder="请输入章节内容" clearable></Input>
                                      </Col>
                                  </Row>
                                  <Row style="margin-top:10px;">
                                      <Col span="18">
                                        <Upload
                                            ref="upload"
                                            :show-upload-list="false"
                                            :default-file-list="defaultList"
                                            :format="['jpg','jpeg','png']"
                                            :max-size="2048"
                                            multiple
                                            type="drag"
                                            action="//jsonplaceholder.typicode.com/posts/"
                                            style="display: inline-block;width:58px;">
                                            <div style="width: 58px;height:58px;line-height: 58px;">
                                                <Icon type="ios-camera" size="20"></Icon>
                                            </div>
                                        </Upload>
                                      </Col>
                                  </Row>
                              </FormItem>
                              <FormItem>
                                  <Row>
                                      <Col span="12">
                                          <Button type="dashed" long @click="addParagh" icon="md-add">Add item</Button>
                                      </Col>
                                  </Row>
                              </FormItem>
                              <FormItem>
                                  <Button type="primary" @click="addBlog('Paragraphs')">保存</Button>
                                  <Button @click="handleReset('Paragraphs')" style="margin-left: 8px">重置</Button>
                              </FormItem>
                          </Form>
                        </Card>
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
                        <Icon type="ios-options" /><span style="margin:5px;">分类：</span>
                    </Row>
                    <Row>
                        <hr  style="height:1px;border:none;border-top:1px solid rgb(136, 135, 135);margin:10px;"/>
                    </Row>
                    <Row>
                        <Tag checkable color="primary">Vue</Tag>
                        <Tag checkable color="success">Java</Tag>
                        <Tag checkable color="error">C#</Tag>
                        <Tag checkable color="warning">Golang</Tag>
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
import HeadNav from '@/components/HeadNav'
    export default {
      name: "add-blog",
      components:{
            HeadNav,
        },
        data () {
            return {
                index: 1,
                Paragraphs: {
                    items: [
                        {
                            Title: '',
                            index: 1,
                            status: 1,
                            Content:"",
                            Image:[]
                        }
                    ]
                },
                formItem: {
                    input: '',
                    select: '',
                    radio: 'male',
                    checkbox: [],
                    switch: true,
                    date: '',
                    time: '',
                    slider: [20, 50],
                    textarea: ''
                },
                types:[],
                selectTypeId:0,
                blog:{
                    TypeId:[],
                    AuthorId:0,
                    Article:{
                        Title:"",
                        Content:"",
                        Image:""
                    },
                    Paragraphs:[]
                }
            }
        },
        methods: {
            addBlog (name) {
                console.log("作者：",localStorage.getItem("userId"))
                this.blog.AuthorId=parseInt(localStorage.getItem("userId"))
                this.blog.Paragraphs=this.Paragraphs.items
                this.blog.TypeId.push(this.selectTypeId)
                console.log(this.blog)
                this.$post("/article", 
                        this.blog
                    ).then(
                    res => {
                        if (res.Code == 200) {
                             this.$Message.success("添加成功");
                        }
                    },
                    err => {
                        console.log(err);
                    }
                );
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        this.$Message.success('Success!');
                    } else {
                        this.$Message.error('Fail!');
                    }
                })
            },
            handleReset (name) {
                this.$refs[name].resetFields();
            },
            addParagh () {
                this.index++;
                this.Paragraphs.items.push({
                    Title: '',
                    index: this.index,
                    status: 1,
                    Content:"",
                    Image:[]
                });
            },
            removeParagh (index) {
                this.Paragraphs.items[index].status = 0;
            },
            getArticleType(){
                this.$fetch("/type",{
                }).then(
                    res => {
                        if (res.Code == 200) {
                            this.types=res.Msg
                        }
                    },
                    err => {
                     console.log(err);
                    }
                );
            }, 
            saveBlog() {
                this.$post("/article", 
                        this.blog
                    ).then(
                    res => {
                        if (res.Code == 200) {
                             this.$Message.success("添加成功");
                        }
                    },
                    err => {
                        console.log(err);
                    }
                );
            }
        },
        mounted(){
            this.getArticleType()
        }
    }
</script>

<style scoped>
.add-blog {
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