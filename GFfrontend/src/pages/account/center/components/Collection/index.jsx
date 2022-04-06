import {
  ContactsOutlined,
  LikeOutlined,
  LikeTwoTone,
  LoadingOutlined,
  MessageOutlined,
  StarOutlined,
  StarTwoTone,
} from '@ant-design/icons';
import { Button, Card, Col, Form, List, Row, Select, Tag, Tabs } from 'antd';
import React from 'react';
import { useRequest, history } from 'umi';
import ArticleListContent from './articleContent/index';
import { getPersonalCollection, removeCollection } from '@/services/user';
import styles from './style.less';

const { Option } = Select;
const FormItem = Form.Item;
const pageSize = 10;
const pageNO = 1;
const username = history.location.search.substring(1);

const Collection = () => {
  const [form] = Form.useForm();
  const { data, reload, loading, loadMore, loadingMore } = useRequest(
    async() => {
      const result = await getPersonalCollection({
        pageSize: pageSize,
        pageNO: pageNO,
      });
      return result;
    },
    {
      loadMore: true,
      formatResult: result => result,
    },
  );

  console.log(data);

  let list = [];
  if(typeof(data.articleDetails)!='undefined') {
    list = data.articleDetails;
  }

  const IconText = ({ type, text, value }) => {
    const icon = {
      type: type,
      text: text,
      value: value,
    };
    switch (type) {
      case 'star-o':
        return (
          <span>
            <StarTwoTone
              style={{
                marginRight: 8,
              }}
              onClick={(e) => onCollection(icon, e)}
            />
            {text}
          </span>
        );
      case 'like-o':
        if(value === false) {
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
        }
        else {
          return (
            <span>
              <LikeTwoTone
                style={{
                  marginRight: 8,
                }}
              />
              {text}
            </span>
          );
        }
        

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

  const onCollection = async(values) => {
    console.log(values);
    const id = values.value;
    const result = await removeCollection({
      username: username,
      postid: id,
    });
    if(result.message === 'Ok') {
      location. reload();
    }
    else {
      
    }
  }

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

  return (
    <>
      <Card
        // style={{
        //   marginTop: 24,
        // }}
        bordered={false}
        // bodyStyle={{
        //   padding: '8px 32px 32px 32px',
        // }}
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
              key={item.ID}
              actions={[
                <IconText key="collection" type="star-o" value={item.Favorited} text={item.NumFavorite} />,
                <IconText key="like" type="like-o" value={item.Liked} text={item.NumLike} />,
                <IconText key="reply" type="message" value={item.ID} text={item.NumComment} />,
              ]}
              //extra={<div className={styles.listItemExtra} />}
            >
              <List.Item.Meta
                title={
                  <a className={styles.listItemMetaTitle} href={'/group/post?'+item.ID}>
                    {item.Title}
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

export default Collection;
