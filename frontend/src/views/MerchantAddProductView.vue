<template>
  <el-container class="add-product-container">
    <el-main>
      <Functions />
      <el-card class="add-product-card">
        <template #header>
          <div class="card-header">
            <h2>添加商品</h2>
          </div>
        </template>

        <el-form ref="addProductFormRef" :model="product" label-width="120px">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="商品名称">
                <el-input v-model="product.name" placeholder="请输入商品名称" />
              </el-form-item>
            </el-col>

            <el-col :span="12">
              <el-form-item label="商品类别">
                <el-select v-model="product.categories" multiple placeholder="请选择商品类别" style="width: 100%">
                  <el-option v-for="category in categories.slice(1)" :key="category.id" :label="category.name"
                    :value="category.id" />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="商品主图">
            <div class="upload-container">
              <el-upload
                class="avatar-uploader"
                action="/api/uploadImage"
                :show-file-list="false"
                :on-success="handleMainUpload"
                :before-upload="beforeAvatarUpload"
              >
                <img v-if="product.imgUrl" :src="product.imgUrl" class="product-image">
                <div v-else class="upload-placeholder">
                  <el-icon class="upload-icon"><Plus /></el-icon>
                  <span>点击上传</span>
                </div>
              </el-upload>
              <div class="upload-tip">
                <el-icon><InfoFilled /></el-icon>
                <div>
                  <div>商品主图</div>
                  <div class="tip-detail">建议尺寸：800x800px，支持JPG/PNG格式</div>
                </div>
              </div>
            </div>
          </el-form-item>

          <el-form-item label="轮播图内容">
            <div class="slider-container">
              <el-row :gutter="20">
                <el-col :span="12" v-for="(item, index) in product.slider_items" :key="index">
                  <div class="slider-item">
                    <div class="media-upload">
                      <el-upload
                        class="slider-uploader"
                        action="/api/uploadImage"
                        :show-file-list="false"
                        :on-success="(res) => handleImageUpload(res, index)"
                        :before-upload="beforeImageUpload"
                      >
                        <img v-if="item.image" :src="item.image" class="slider-image" />
                        <div v-else class="upload-placeholder">
                          <el-icon><Plus /></el-icon>
                          <div>上传图片</div>
                        </div>
                      </el-upload>

                      <el-upload
                        class="slider-uploader"
                        action="/api/uploadImage"
                        :show-file-list="false"
                        :on-success="(res) => handleVideoUpload(res, index)"
                        :before-upload="beforeVideoUpload"
                      >
                        <video v-if="item.video" :src="item.video" class="slider-video" controls />
                        <div v-else class="upload-placeholder">
                          <el-icon><VideoCamera /></el-icon>
                          <div>上传视频</div>
                        </div>
                      </el-upload>

                      <el-button type="danger" circle class="remove-button" @click="removeCombination(index)">
                        <el-icon><Delete /></el-icon>
                      </el-button>
                    </div>
                  </div>
                </el-col>
              </el-row>

              <el-button
                v-if="product.slider_items.length < 5"
                class="add-combination"
                @click="addCombination"
              >
                <el-icon><Plus /></el-icon>
                添加内容组合（剩余 {{ 5 - product.slider_items.length }} 个）
              </el-button>
            </div>
          </el-form-item>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="商品价格">
                <el-input-number
                  v-model="product.price"
                  :precision="2"
                  :step="0.1"
                  :min="0"
                  style="width: 100%"
                  placeholder="请输入商品价格"
                />
              </el-form-item>
            </el-col>

            <el-col :span="12">
              <el-form-item label="商品库存">
                <el-input-number
                  v-model="product.stock"
                  :min="0"
                  :step="1"
                  style="width: 100%"
                  placeholder="请输入商品库存"
                />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item label="商品描述">
            <el-input
              v-model="product.description"
              type="textarea"
              :rows="4"
              placeholder="请输入商品描述"
            />
          </el-form-item>

          <el-form-item>
            <div class="form-buttons">
              <el-button @click="resetForm">重置</el-button>
              <el-button type="primary" @click="addProduct">添加商品</el-button>
            </div>
          </el-form-item>
        </el-form>
      </el-card>
    </el-main>
  </el-container>
</template>

<script setup>
import Functions from '@/components/merchant/Functions.vue';
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { ElMessage } from 'element-plus';
import axios from 'axios';

const router = useRouter();
const store = useStore();
const addProductFormRef = ref(null);

// Vuex 数据
const categories = computed(() => store.getters['category/categories']);


const product = ref({
  name: '',
  categories: [],
  price: 0,
  stock: 0,
  description: '',
  imgUrl: '',
  // 修复初始数据结构不一致问题
  slider_items: []  // 替换原来的 slider_imgs: [[]]
});


const resetForm = () => {
  product.value = {
    name: '',
    categories: [],
    price: 0,
    stock: 0,
    description: '',
    imgUrl: '',
    slider_items: []  // 保持数据结构一致
  };
  activeSliderIndex.value = 0;
};


const isVideo = (url) => {
  return url?.match(/\.(mp4|mov|avi)$/i);
};

const addProduct = async () => {
  try {
    console.log('ProductDetail:', product)
    const payload = {
      ...product.value,
      category: product.value.categories.map(id => ({
        id,
        name: categories.value.find(c => c.id === id)?.name
      }))
    };

    await axios.post('/api/merchant/product/add', payload);
    ElMessage.success('商品添加成功');
    router.push('/merchant/products');
  } catch (error) {
    ElMessage.error('商品添加失败：' + error.message);
  }
};


const beforeAvatarUpload = (file) => {
  const isImage = ['image/jpeg', 'image/png'].includes(file.type);
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isImage) {
    ElMessage.error('只支持 JPG/PNG 格式!');
    return false;
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!');
    return false;
  }
  return true;
};

// 在 script setup 中添加缺失的方法
const removeSlider = (index) => {
  product.value.slider_imgs.splice(index, 1)
}

const handleSliderUpload = (response) => {
  if (response.url) {
    if (product.value.slider_imgs.length >= 5) {
      ElMessage.warning('最多只能上传5个内容')
      return
    }
    product.value.slider_imgs.push(response.url)
  }
}

const beforeSliderUpload = () => {
  if (product.value.slider_imgs.length >= 5) {
    ElMessage.error('轮播图最多只能上传5个内容')
    return false
  }
  return true
}

// 修正后的 handleMainUpload 方法
const handleMainUpload = (response) => {
  if (response.url) {
    product.value.imgUrl = response.url
    // 自动插入到轮播图首位并去重
    product.value.slider_imgs = [
      response.url,
      ...product.value.slider_imgs.filter(url => url !== response.url)
    ].slice(0, 5)
    ElMessage.success('主图上传成功')
  } else {
    ElMessage.error('上传失败')
  }
}

const imageList = computed(() => product.value.slider_imgs.filter(item => !isVideo(item)));
const videoList = computed(() => product.value.slider_imgs.filter(item => isVideo(item)));


const beforeImageUpload = (file) => {
  const isImage = ['image/jpeg', 'image/png'].includes(file.type);
  if (!isImage) {
    ElMessage.error('只支持 JPG/PNG 格式!');
    return false;
  }
  return true;
};

const beforeVideoUpload = (file) => {
  const isVideo = file.type === 'video/mp4';
  const isLt50M = file.size / 1024 / 1024 < 50;

  if (!isVideo) {
    ElMessage.error('只支持 MP4 格式!');
    return false;
  }
  if (!isLt50M) {
    ElMessage.error('视频大小不能超过 50MB!');
    return false;
  }
  return true;
};

// 添加组合条目方法
const addCombination = () => {
  if (product.value.slider_items.length < 5) {
    product.value.slider_items.push({ image: null, video: null });
  }
};

// 修改后的上传处理方法
const handleImageUpload = (response, index) => {
  if (response.url) {
    product.value.slider_items[index].image = response.url;
  }
};

const handleVideoUpload = (response, index) => {
  if (response.url) {
    product.value.slider_items[index].video = response.url;
  }
};

// 删除媒体内容方法
const removeMedia = (index, type) => {
  product.value.slider_items[index][type] = null;
};

// 删除整个组合
const removeCombination = (index) => {
  product.value.slider_items.splice(index, 1);
};
</script>

<style scoped>
.add-product-container {
  padding: 20px;
  background-color: #f5f7fa;
  min-height: calc(100vh - 60px);
}

.add-product-card {
  max-width: 1200px;
  margin: 20px auto;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0,0,0,0.1);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h2 {
  margin: 0;
  font-size: 20px;
  color: var(--el-text-color-primary);
  font-weight: 600;
}

.upload-container {
  display: flex;
  align-items: flex-start;
  gap: 20px;
}

.product-image {
  width: 200px;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
}

.upload-placeholder {
  width: 200px;
  height: 200px;
  border: 1px dashed var(--el-border-color);
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: var(--el-text-color-placeholder);
  cursor: pointer;
  transition: all 0.3s;
}

.upload-placeholder:hover {
  border-color: var(--el-color-primary);
  color: var(--el-color-primary);
}

.upload-icon {
  font-size: 28px;
  margin-bottom: 8px;
}

.upload-tip {
  display: flex;
  align-items: flex-start;
  gap: 8px;
  color: var(--el-text-color-secondary);
  font-size: 14px;
}

.tip-detail {
  margin-top: 4px;
  font-size: 12px;
  color: var(--el-text-color-placeholder);
}

.slider-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.slider-item {
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  position: relative;
  margin-bottom: 20px;
}

.media-upload {
  display: flex;
  gap: 20px;
  position: relative;
}

.slider-image, .slider-video {
  width: 200px;
  height: 200px;
  object-fit: cover;
  border-radius: 8px;
}

.remove-button {
  position: absolute;
  right: -10px;
  top: -10px;
  z-index: 1;
}

.add-combination {
  padding: 30px;
  border: 2px dashed var(--el-border-color-lighter);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  color: var(--el-text-color-secondary);
  transition: all 0.3s;
  background: white;
}

.add-combination:hover {
  border-color: var(--el-color-primary);
  color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
}

.form-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 20px;
}

:deep(.el-form-item__label) {
  font-weight: 500;
}

:deep(.el-input-number) {
  width: 100%;
}

/* 响应式布局 */
@media (max-width: 768px) {
  .add-product-card {
    margin: 10px;
  }

  .upload-container {
    flex-direction: column;
  }

  .media-upload {
    flex-direction: column;
  }

  .product-image,
  .slider-image,
  .slider-video,
  .upload-placeholder {
    width: 100%;
  }
}
</style>