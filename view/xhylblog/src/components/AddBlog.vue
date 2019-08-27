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
                                        <Input v-model="formItem.input" placeholder="Title" clearable></Input>
                                    </FormItem>
                                    <FormItem label="类型：" style="width: 50%;">
                                        <Select v-model="blog.TypeId" placeholder="请选择类型">
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
                                        <Input v-model="formItem.textarea" type="textarea" :autosize="{minRows: 2,maxRows: 5}" placeholder="Enter something..." clearable></Input>
                                    </FormItem>
                                </Form>
                            </div>
                        </Card>
                        <Card class="card container">
                            <Form ref="formDynamic" :model="formDynamic" :label-width="80" style="width: 100;">
                              <FormItem
                                      v-for="(item, index) in formDynamic.items"
                                      v-if="item.status"
                                      :key="index"
                                      :label="'章节 ' + item.index"
                                      :prop="'items.' + index + '.value'"
                                      :rules="{required: true, message: 'Item ' + item.index +' can not be empty', trigger: 'blur'}">
                                  <Row>
                                      <Col span="18">
                                          <Input type="text" v-model="item.value" placeholder="请输入章节标题" clearable></Input>
                                      </Col>
                                      <Col span="4" offset="1">
                                          <Button @click="handleRemove(index)" type="error">删除</Button>
                                      </Col>
                                  </Row>
                                  <Row>
                                      <Col span="18">
                                          <Input v-model="pcontent" type="textarea" :autosize="{minRows: 5,maxRows: 8}" placeholder="请输入章节内容" clearable></Input>
                                      </Col>
                                  </Row>
                                  <Row>
                                  </Row>
                              </FormItem>
                              <FormItem>
                                  <Row>
                                      <Col span="12">
                                          <Button type="dashed" long @click="handleAdd" icon="md-add">Add item</Button>
                                      </Col>
                                  </Row>
                              </FormItem>
                              <FormItem>
                                  <Button type="primary" @click="handleSubmit('formDynamic')">保存</Button>
                                  <Button @click="handleReset('formDynamic')" style="margin-left: 8px">重置</Button>
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
                formDynamic: {
                    items: [
                        {
                            value: '',
                            index: 1,
                            status: 1
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
                blog:{
                    TypeId:0,
                    Article:{
                        Title:"",
                        Content:"Content",
                        Image:"a.png"
                    },
                    Paragraphs:[]
                }
            }
        },
        methods: {
            handleSubmit (name) {
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
            handleAdd () {
                this.index++;
                this.formDynamic.items.push({
                    value: '',
                    index: this.index,
                    status: 1
                });
            },
            handleRemove (index) {
                this.formDynamic.items[index].status = 0;
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
                            alert("添加成功");
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