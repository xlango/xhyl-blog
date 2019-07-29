<template>
<div class="add-blog">
<Form ref="formValidate" :model="blog" :rules="blogrule" :label-width="80">
        <FormItem label="标题：" prop="title">
            <Input v-model="blog.title" placeholder="请输入标题"></Input>
        </FormItem>
        <FormItem label="作者：" prop="author">
            <Select v-model="blog.author" placeholder="请选择作者">
                <Option v-for="author in authors" :key="author.id">
                    {{author.name}}
                </Option>
            </Select>
        </FormItem>
        <FormItem label="分类：" prop="type">
            <CheckboxGroup v-model="blog.types">
                <Checkbox :label="item.name" 
                    v-for="item in types" v-bind:key="item.id">
                    {{item.name}}
                </Checkbox>
            </CheckboxGroup>
        </FormItem>
        <FormItem label="内容：" prop="content">
            <Input v-model="blog.content" type="textarea" :autosize="{minRows: 5,maxRows: 5}" placeholder="请输入内容"></Input>
        </FormItem>
        <FormItem>
            <Button type="primary" >保存</Button>
            <Button style="margin-left: 8px">重置</Button>
        </FormItem>
    </Form>
    <ul>
        <li v-for="(blogtype,index) in blog.types" 
            :key="index">
            {{blogtype}}
        </li>
    </ul>
</div> 
    
</template>
<script>
    export default {
        name:"add-blog",
        data () {
             return {
                blog: {
                    title: '',
                    author: '',
                    types: [],
                    content: ''
                },
                blogrule: {
                    title: [
                        { required: true, message: '标题不能为空', trigger: 'blur' }
                    ],
                    author: [
                        { required: true, message: '请选择作者', trigger: 'blur' }
                    ],
                    types: [
                        { required: true, type: 'array', min: 1, message: '至少选择一个类型', trigger: 'blur' },
                        { type: 'array', max: 2, message: '请选择类型', trigger: 'blur' }
                    ],
                    content: [
                        { required: true, message: '请输入内容', trigger: 'blur' },
                        { type: 'string', min: 20, message: '不少于20字符', trigger: 'blur' }
                    ]
                },
                authors: [
                        {
                            id:1,
                            name:"张三"
                        },
                        {
                            id:2,
                            name:"李四"
                        },
                        {
                            id:3,
                            name:"王五"
                        }],
                types:[
                    {
                            id:1,
                            name:"Java"
                        },
                        {
                            id:2,
                            name:"Golang"
                        },
                        {
                            id:3,
                            name:"C#"
                        },
                        {
                            id:4,
                            name:"Vue"
                        }
                ]
            }
        },
        methods: {

        }
    }
</script>

<style scoped>
body{
    text-align: center;
}
.add-blog{ margin:0 auto; width:400px; height:100px;}
</style>
