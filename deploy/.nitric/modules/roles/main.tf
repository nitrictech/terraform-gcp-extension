# TODO: Override nitric provided roles with known pre-deployed roles
locals {
  bucket_reader_role = ""
  bucket_writer_role = ""
  bucket_deleter_role = ""
  topic_publisher_role = ""
  base_compute_role = ""
  secret_access_role = ""
  secret_put_role = ""
  kv_reader_role = ""
  kv_writer_role = ""
  kv_deleter_role = ""
  queue_enqueue_role = ""
  queue_dequeue_role = ""
}
