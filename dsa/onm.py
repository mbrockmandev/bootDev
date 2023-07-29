import random


def get_avg_brand_followers(all_handles, brand_name):
    count = 0
    for i in range(len(all_handles)):
        for j in range(len(all_handles[i])):
            if brand_name in all_handles[i][j]:
                count += 1
    return count / len(all_handles)


# don't touch below this line


def test(num_handles, avg_aud_size, brand_name):
    print(
        f"Checking {num_handles} influencers with average audience sizes of {avg_aud_size}..."
    )
    all_handles = get_all_handles(num_handles, avg_aud_size)
    avg = round(get_avg_brand_followers(all_handles, brand_name), 2)
    print(f"Average {brand_name} fans per influencer: {avg}")
    print("====================================")


def main():
    random.seed(1)
    test(10, 1000, "luxa")
    test(20, 2000, "luxa")
    test(40, 4000, "luxa")
    test(80, 8000, "luxa")
    test(160, 16000, "luxa")


def get_all_handles(num, audience_size):
    all_handles = []
    for i in range(num):
        m = random.randrange(
            int(audience_size - audience_size * 1.2),
            int(audience_size + audience_size * 1.2),
        )
        handles = get_instagram_handles(m)
        all_handles.append(handles)
    return all_handles


def get_instagram_handles(num):
    handles = []
    for i in range(0, num):
        m = random.randrange(0, 6)
        if m == 0:
            handles.append(f"luxaraygirl{i}")
        elif m == 1:
            handles.append(f"theprimerog{i}")
        elif m == 2:
            handles.append(f"luxafanboi{i}")
        elif m == 3:
            handles.append(f"dashlord{i}")
        elif m == 4:
            handles.append(f"saintrex{i}")
        elif m == 5:
            handles.append(f"writergurl{i}")
    return handles


main()
