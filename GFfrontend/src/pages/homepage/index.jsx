import { LikeOutlined, LoadingOutlined, MessageOutlined, StarOutlined } from '@ant-design/icons';
import { Button, Card, Col, Form, List, Row, Select, Tag } from 'antd';
import React, { useState } from 'react';
import { useRequest, useModel, history } from 'umi';
import ArticleListContent from './components/ArticleListContent';
import StandardFormRow from './components/StandardFormRow';
import TagSelect from './components/TagSelect';
import { queryList } from '@/services/getList';
import { createLike } from '@/services/user';
import styles from './style.less';

const { Option } = Select;
const FormItem = Form.Item;
const pageSize = 20;
const pageNumber = 1;


const Articles = () => {
  const [form] = Form.useForm();
  const { initialState } = useModel('@@initialState');
  const { currentUser } = initialState || {};

  const { data, reload, loading, loadMore, loadingMore } = useRequest(
    async() => {
      const result = await queryList({
        PageNO: pageNumber,
        PageSize: pageSize,
      });
      return result;
    },
    {
      formatResult: result => result,
      loadMore: true,
    }    
  );

  //console.log(data);
  const list = [];
  if(typeof(data.ArticleList)!='undefined') {
    const articleList = data.ArticleList;
    const communityList = data.CommunityList;
    const collection = data.CountFavorite;
    const like = data.CountLike;
    const reply = data.CountComment;
    const size = Object.keys(articleList).length;
    for(let i=0; i<size; i++) {
      list.push({
        id: articleList[i].ID,
        name: articleList[i].Username,
        title: articleList[i].Title,
        group: communityList[i].Name,
        createdAt: articleList[i].CreateDay,
        content: articleList[i].Content,
        collection: collection[i],
        like: like[i],
        reply: reply[i],
        groupID: communityList[i].ID,
        avatar: 'http://10.20.0.166:10010/resources/userfiles/'+ articleList[i].Username+'/avatar.png',
      });
    }
  }
  console.log(list);

  const onCCollection = async(values) => {
    console.log(values);
    if(values.type === 'star-o') {
      if(values.value === '1') {
        return (
          <IconText key="collection" type="star-o" value="0" text={values.text--} />
        );
      }
      else {
        return (
          <IconText key="collection" type="star-o" value="1" text={item.collection} />
        )
      }
    }
  }

  const onLike = async(values) => {
    console.log("liked");
    console.log(values);
    const result = await createLike({
      id: values,
    });
    if(result === '200') {
      location.reload();
    }
  }

  const IconText = ({ value, type, text }) => {
    switch (type) {
      case 'star-o':
        return (
          <span>
            <StarOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      case 'like-o':
        return (
          <span>
            <LikeOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      case 'message':
        return (
          <span>
            <MessageOutlined
              style={{
                marginRight: 8,
              }}
            />
            {text}
          </span>
        );

      default:
        return null;
    }
  };

  const formItemLayout = {
    wrapperCol: {
      xs: {
        span: 24,
      },
      sm: {
        span: 24,
      },
      md: {
        span: 12,
      },
    },
  };

  const onCollection = async(values) => {
    console.log(values);
    let count = values;
    count++;
    return count;
  }

  const loadMoreDom = list.length > 0 && (
    <div
      style={{
        textAlign: 'center',
        marginTop: 16,
      }}
    >
      <Button
        onClick={loadMore}
        style={{
          paddingLeft: 48,
          paddingRight: 48,
        }}
      >
        {loadingMore ? (
          <span>
            <LoadingOutlined /> Loading...
          </span>
        ) : (
          'Load More'
        )}
      </Button>
    </div>
  );

  const clickPost = (values) => {
    history.push({
      pathname: '/group/post',
      search: values.toString(),
    });
    return;
  }

  return (
    <>
      <Card bordered={false}>
        <Form
          layout="inline"
          form={form}
          initialValues={
            {
              //owner: ['wjh', 'zxx'],
            }
          }
          onValuesChange={reload}
        >
        <TagSelect expandable>
          <TagSelect.Option value="cat1">Sports</TagSelect.Option>
          <TagSelect.Option value="cat2">Professors</TagSelect.Option>
          <TagSelect.Option value="cat3">Courses</TagSelect.Option>
          <TagSelect.Option value="cat4">Daily Life</TagSelect.Option>
          <TagSelect.Option value="cat5">Movies</TagSelect.Option>
        </TagSelect>
        </Form>
      </Card>
      <Card
        style={{
          marginTop: 10,
        }}
        bordered={false}
        bodyStyle={{
          padding: '8px 32px 32px 32px',
        }}
      >
        <List
          size="large"
          loading={loading}
          rowKey="id"
          itemLayout="vertical"
          loadMore={loadMoreDom}
          dataSource={list}
          renderItem={(item) => (
            <List.Item
              key={item.id}
              actions={[
                <IconText key="collection" type="star-o" value = {item.id} text={item.collection}  />,
                <IconText key="like" type="like-o" value = {item.id} text={item.like} />,
                <IconText key="reply" type="message" value = {item.id} text={item.reply} />,
              ]}
            >
              <List.Item.Meta
                title={
                  <a className={styles.listItemMetaTitle}  onClick={(e) => clickPost(item.id, e)}>
                    {item.title}
                  </a>
                }
              />
              <ArticleListContent data={item} />
            </List.Item>
          )}
        />
      </Card>
    </>
  );
};

export default Articles;
